package observer

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	diverclient "github.com/VigorousDeveloper/poc-human/processor/humanclient"
	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/cenkalti/backoff"
	stypes "github.com/cosmos/cosmos-sdk/types"
)

// Observer observer service
type Observer struct {
	lock             *sync.Mutex
	stopChan         chan struct{}
	HumanChainBridge *diverclient.HumanChainBridge
	storage          *ObserverStorage

	CurEthHeight   uint64
	CurHumanHeight uint64

	approve_voted []string
	keysign_voted []string

	EthPoolChanged chan bool
	HmPoolChanged  chan bool

	ArrMsgUpdateBalance      []*types.MsgUpdateBalance
	ArrMsgKeysignVote        []*types.MsgKeysignVote
	ArrMsgObservationVote    []*types.MsgObservationVote
	ArrMsgApproveTransaction []*types.MsgApproveTransaction
	ArrMsgTranfserPoolcoin   []*types.MsgTranfserPoolcoin

	ArrMsgsToSend []*stypes.Msg
}

const (
	//----------ETHEREUM---------
	//-------------------------
	// Ethereum RPC Node Provider URL from Alchemy
	URL_Ethereum_RPC_Node_Provider = "https://eth-rinkeby.alchemyapi.io/v2/JiLlXSz2HgdpuutVqt-irguqqDBxPV4D"

	// Ethereum RPC Node Provider WSS URL from Alchemy, rinkeby
	URL_Ethereum_RPC_Node_Provider_WSS = "wss://eth-rinkeby.alchemyapi.io/v2/JiLlXSz2HgdpuutVqt-irguqqDBxPV4D"

	// Ethereum Rinkeby USDK Contract Address
	Ethereum_USDK_Token_Address = "0x7Ba1E70BF249eEF06de34af3E2695eFccFc4a0d2"

	// Ethereum Pool Account Address
	Ethereum_Pool_Address = "0x369b28f227C0188478cb05F8467bdd52002EcC4E"

	// Ethereum Pool Account Private Key
	Ethereum_Pool_Account_Private_Key = "4b11634f979c262e33def94f52a0a82e57d0db5d7f94efd2844a1892623e063c"
)

// NewObserver create a new instance of Observer for chain
func NewObserver(chainBridge *diverclient.HumanChainBridge, dataPath string) (*Observer, error) {
	storage, err := NewObserverStorage(dataPath)
	if err != nil {
		return nil, fmt.Errorf("fail to create observer storage: %w", err)
	}

	return &Observer{
		lock:                  &sync.Mutex{},
		stopChan:              make(chan struct{}),
		HumanChainBridge:      chainBridge,
		storage:               storage,
		CurEthHeight:          0,
		CurHumanHeight:        0,
		EthPoolChanged:        make(chan bool),
		HmPoolChanged:         make(chan bool),
		approve_voted:         make([]string, 0),
		keysign_voted:         make([]string, 0),
		ArrMsgUpdateBalance:   make([]*types.MsgUpdateBalance, 0),
		ArrMsgKeysignVote:     make([]*types.MsgKeysignVote, 0),
		ArrMsgObservationVote: make([]*types.MsgObservationVote, 0),
	}, nil
}

// Check Ethereum Block Count through Gorelli RPC provider
func (o *Observer) DoCurRequest(payload io.Reader, endpoint string) string {
	// generate POST request to the Solana Node RPC provider
	req, _ := http.NewRequest("POST", endpoint, payload)

	// Content-type to application/json
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

func (o *Observer) Start() error {
	go o.ProcessTxInsEthExternal()
	go o.ProcessTxInsHmExternal()
	go o.ProcessUpdateEthPoolBalance()
	go o.ProcessUpdateHumanPoolBalance()
	go o.ProcessKeysignTx()
	go o.ProcessSendTxToHumanChain()

	o.EthPoolChanged <- true
	o.HmPoolChanged <- true

	return nil
}

func (o *Observer) Stop() error {
	close(o.stopChan)

	return nil
}

// Eth Balance Checking
func (o *Observer) ProcessUpdateEthPoolBalance() {
	for {
		select {
		case <-o.stopChan:
			return
		case <-o.EthPoolChanged:
			o.FetchBalanceOfEtherPool()
		}
	}
}

// Human Balance Checking
func (o *Observer) ProcessUpdateHumanPoolBalance() {
	for {
		select {
		case <-o.stopChan:
			return
		case <-o.HmPoolChanged:
			o.FetchBalanceOfEtherPool()
		}
	}
}

func (o *Observer) ProcessKeysignTx() {
	for {
		select {
		case <-o.stopChan:
			return
		case <-time.After(time.Second * 2):
			o.FetchTransactionAndBroadcastKeysignTx()
		}
	}
}

// Ethereum Checking
func (o *Observer) ProcessSendTxToHumanChain() {
	for {
		select {
		case <-o.stopChan:
			return
		case <-time.After(time.Second):
			o.SendTxToDiversifiChain()
		}
	}
}

// Send msgs to diversifi chain
func (o *Observer) SendTxToDiversifiChain() error {
	msgs := make([]stypes.Msg, 0)
	for _, m := range o.ArrMsgKeysignVote {
		msgs = append(msgs, m)
	}

	for _, m := range o.ArrMsgObservationVote {
		msgs = append(msgs, m)
	}

	for _, m := range o.ArrMsgUpdateBalance {
		msgs = append(msgs, m)
	}

	for _, m := range o.ArrMsgApproveTransaction {
		msgs = append(msgs, m)
	}

	for _, m := range o.ArrMsgTranfserPoolcoin {
		msgs = append(msgs, m)
	}

	if len(msgs) < 1 {
		return nil
	}

	err := o.SendBroadcast(msgs...)
	if err == nil {
		o.ArrMsgKeysignVote = o.ArrMsgKeysignVote[:0]
		o.ArrMsgObservationVote = o.ArrMsgObservationVote[:0]
		o.ArrMsgUpdateBalance = o.ArrMsgUpdateBalance[:0]
		o.ArrMsgApproveTransaction = o.ArrMsgApproveTransaction[:0]
		o.ArrMsgTranfserPoolcoin = o.ArrMsgTranfserPoolcoin[:0]
	}

	return err
}

// Send try b
func (o *Observer) SendBroadcast(msgs ...stypes.Msg) error {
	bf := backoff.NewExponentialBackOff()
	bf.MaxElapsedTime = time.Second * 5

	return backoff.Retry(func() error {
		_, err := o.HumanChainBridge.Broadcast(msgs...)
		if err != nil {
			return fmt.Errorf("fail to send the tx to thorchain: %w", err)
		}

		return nil
	}, bf)
}

// check if a slice continas a string
func (o *Observer) continsHash(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// Fetch DiversifiChain & Broadcast Keysign Transaction
func (o *Observer) FetchTransactionAndBroadcastKeysignTx() bool {
	// Get PubKey & Voter Address
	pubKey, voter := o.HumanChainBridge.GetVoterInfo()

	// Get All Transaction Data from DiversifiChain
	txDataList, err := o.HumanChainBridge.GetTxDataList("")

	//
	if err != nil {
		return false
	}

	// Looping
	for _, tx := range txDataList.TransactionData {
		// If it is not confirmed, continue
		if tx.Status != types.PAY_CONFIRMED && tx.Status != types.PAY_KEYSIGNED {
			continue
		}

		// It shouldn't be pay confirmed and voted hash.
		if tx.Status == types.PAY_CONFIRMED && !o.continsHash(o.keysign_voted, tx.ConfirmedBlockHash) {
			// observe voted list
			o.keysign_voted = append(o.keysign_voted, tx.ConfirmedBlockHash)

			// construct msg
			msg := types.NewMsgKeysignVote(voter, tx.ConfirmedBlockHash, pubKey)
			o.ArrMsgKeysignVote = append(o.ArrMsgKeysignVote, msg)

			if tx.OriginChain == types.CHAIN_ETHEREUM {
				o.EthPoolChanged <- true
			}

			if tx.OriginChain == types.CHAIN_HUMAN {
				o.HmPoolChanged <- true
			}

			return true
		}

		// It shouldn't be pay confirmed and voted hash.
		if tx.Status == types.PAY_KEYSIGNED && !o.continsHash(o.approve_voted, tx.ConfirmedBlockHash) {
			// observe voted list
			o.approve_voted = append(o.approve_voted, tx.ConfirmedBlockHash)

			moniker := o.HumanChainBridge.GetMonikerName()
			data := &types.TransactionData{
				Index:              "1",
				OriginChain:        tx.OriginChain,
				OriginAddress:      tx.OriginAddress,
				TargetChain:        tx.TargetChain,
				TargetAddress:      tx.TargetAddress,
				Amount:             tx.Amount,
				Time:               tx.Time,
				Creator:            tx.Creator,
				Status:             tx.Status,
				ConfirmedBlockHash: tx.ConfirmedBlockHash,
				SignedKey:          "",
				Fee:                tx.Fee,
			}

			bResult := false
			if tx.TargetChain == types.CHAIN_ETHEREUM {
				bResult = o.EthereumTransferTokenToTarget(data, moniker)
			} else if tx.TargetChain == types.CHAIN_HUMAN {
				bResult = o.HumanTransferTokenToTarget(data, moniker)
			}

			pubKey, _ := o.HumanChainBridge.GetVoterInfo()
			securedKey := types.EncryptMsgSHA256(pubKey)

			// construct msg
			if bResult {
				msg := types.NewMsgApproveTransaction(voter, tx.ConfirmedBlockHash, types.PAY_PAID, securedKey)
				o.ArrMsgApproveTransaction = append(o.ArrMsgApproveTransaction, msg)
			} else {
				msg := types.NewMsgApproveTransaction(voter, tx.ConfirmedBlockHash, types.PAY_FAILED, securedKey)
				o.ArrMsgApproveTransaction = append(o.ArrMsgApproveTransaction, msg)
			}

			return true
		}
	}

	return true
}
