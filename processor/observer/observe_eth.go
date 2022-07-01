package observer

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	token "github.com/VigorousDeveloper/poc-human/x/pochuman/types/erc20"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	etherTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Fetches the USDC balance of Ethereum pool account
func (o *Observer) FetchBalanceOfEtherPool() bool {
	client, err := ethclient.Dial(URL_Ethereum_RPC_Node_Provider)
	if err != nil {
		return false
	}

	// Golem (GNT) Address
	tokenAddress := common.HexToAddress(Ethereum_USDK_Token_Address)
	instance, err := token.NewMain(tokenAddress, client)
	if err != nil {
		return false
	}

	address := common.HexToAddress(Ethereum_Pool_Address)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return false
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return false
	}

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	_, voter := o.HumanChainBridge.GetVoterInfo()

	msg := types.NewMsgUpdateBalance(voter, types.CHAIN_ETHEREUM, fmt.Sprintf("%f", value), fmt.Sprintf("%v", decimals))
	o.ArrMsgUpdateBalance = append(o.ArrMsgUpdateBalance, msg)

	return true
}

// Transfer token on Ethereum
func (o *Observer) EthereumTransferTokenToTarget(txdata *types.TransactionData, moniker string) bool {
	client, err := ethclient.Dial(URL_Ethereum_RPC_Node_Provider)
	if err != nil {
		fmt.Println(err)
		return false
	}

	privateKey, err := crypto.HexToECDSA(Ethereum_Pool_Account_Private_Key)
	if err != nil {
		fmt.Println(err)
		return false
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return false
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err)
		return false
	}

	value := big.NewInt(0) // in wei (0 eth)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	_, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println(err)
		return false
	}

	toAddress := common.HexToAddress(txdata.TargetAddress)
	tokenAddress := common.HexToAddress(Ethereum_USDK_Token_Address)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

	amt, _ := strconv.ParseFloat(txdata.Amount, 64)

	// -----
	amtFee, err := strconv.ParseFloat(txdata.Fee, 64)
	amt -= amtFee

	samt := fmt.Sprintf("%f", amt*1e18) /// consider the denomination number to 18

	amount := new(big.Int)
	amount.SetString(samt, 10) // sets the value to 1 USDC, in the token denomination

	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})

	if err != nil {
		return false
	}

	tx := etherTypes.NewTransaction(nonce, tokenAddress, value, gasLimit*2, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
		return false
	}

	signedTx, err := etherTypes.SignTx(tx, etherTypes.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// Semaphore for transfer
	if moniker != types.MAIN_VALIDATOR_MONIKER {
		return true
	}

	// Broadcast tx
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
		return false
	}

	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
	return true
}

// Keep listening to WSS and fetch transaction deposited to the pool
func (o *Observer) ProcessTxInsEthExternal() {
	client, err := ethclient.Dial(URL_Ethereum_RPC_Node_Provider_WSS)
	if err != nil {
		return
	}

	contractAddress := common.HexToAddress(Ethereum_USDK_Token_Address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan etherTypes.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		return
	}

	for {
		select {
		case <-o.stopChan:
			return
		case <-sub.Err():
			time.Sleep(time.Second * 5)
		case vLog := <-logs:
			o.EthereumParseLog(vLog)
		}
	}
}

func (o *Observer) EthereumParseLog(vLog etherTypes.Log) {
	contractAbi, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return
	}

	logTransferSig := []byte("Transfer(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	switch vLog.Topics[0].Hex() {
	case logTransferSigHash.Hex():
		var transferEvent LogTransfer

		i, err := contractAbi.Unpack("Transfer", vLog.Data)
		if err != nil {
			return
		}

		amt := i[0].(*big.Int)
		famt := new(big.Float).SetInt(amt)
		ff, _ := famt.Float64()
		tokenAmount := ff / 1e18

		transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
		transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

		if transferEvent.From.String() == Ethereum_Pool_Address {
			// Send true to SolPoolchange channel
			o.EthPoolChanged <- true
			return
		}

		if transferEvent.To.String() != Ethereum_Pool_Address || tokenAmount == 0.0 {
			return
		}

		if transferEvent.From.String() == Ethereum_USDK_Token_Address {
			return
		}

		_, voter := o.HumanChainBridge.GetVoterInfo()
		msg := types.NewMsgObservationVote(voter, vLog.TxHash.String(), types.CHAIN_ETHEREUM, transferEvent.From.Hex(), transferEvent.To.Hex(), fmt.Sprintf("%f", tokenAmount))
		o.ArrMsgObservationVote = append(o.ArrMsgObservationVote, msg)

		// Send true to EthPoolchange channel
		o.EthPoolChanged <- true
	}
}
