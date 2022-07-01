package observer

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	// "github.com/tendermint/tendermint/rpc/coretypes"
)

// Fetches the USDC balance of Ethereum pool account
func (o *Observer) FetchBalanceOfHumanPool() bool {
	accBalance, err := o.HumanChainBridge.GetBalance(types.Humanchain_Pool_Address)
	if err != nil {
		return false
	}

	_, voter := o.HumanChainBridge.GetVoterInfo()

	// Constrcut msg to be broadcasted
	msg := types.NewMsgUpdateBalance(voter, types.CHAIN_HUMAN, fmt.Sprintf("%f", accBalance.Balances[0].Amount/1e9), fmt.Sprintf("%v", 9))
	o.ArrMsgUpdateBalance = append(o.ArrMsgUpdateBalance, msg)

	return true
}

// Transfer token on Human
func (o *Observer) HumanTransferTokenToTarget(txdata *types.TransactionData, moniker string) bool {
	// Semaphore for transfer
	if moniker != types.MAIN_VALIDATOR_MONIKER {
		return true
	}

	_, creator := o.HumanChainBridge.GetVoterInfo()

	// Constrcut 100 uHMN, decimal 9
	famt, err := strconv.ParseFloat(txdata.Amount, 64)
	if err != nil {
		return false
	}

	// Fee
	amtFee, err := strconv.ParseFloat(txdata.Fee, 64)
	famt -= amtFee

	// String conv
	amt := fmt.Sprintf("%fuHMN", famt*1e9)

	// Construct a message to be broadcasted
	msg := types.NewMsgTranfserPoolcoin(creator, txdata.TargetAddress, amt)
	o.ArrMsgTranfserPoolcoin = append(o.ArrMsgTranfserPoolcoin, msg)

	return true
}

// Keep listening to WSS and fetch transaction deposited to the pool
func (o *Observer) ProcessTxInsHmExternal() {
	ctx := o.HumanChainBridge.GetContext()
	client := ctx.Client
	err := client.Start()
	if err != nil {
		return
	}
	defer client.Stop()

	ctx0, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	query := "tm.event = 'Tx'"
	txs, err := ctx.Client.Subscribe(ctx0, "test-client", query)
	if err != nil {
		return
	}

	for {
		select {
		case <-o.stopChan:
			return
		case tx := <-txs:
			o.HumanParseLog(tx.Events)
		}
	}
}

// ---------------------------------- //
// --- Parse Human chain Transfer log //
// ---------------------------------- //
func (o *Observer) HumanParseLog(txs map[string][]string) {
	msgActions := txs["message.action"]
	if len(msgActions) < 1 {
		return
	}
	msgAction := msgActions[0]
	if msgAction != "/cosmos.bank.v1beta1.MsgSend" {
		return
	}

	sender := txs["coin_spent.spender"][0]
	receiver := txs["coin_received.receiver"][0]

	if sender == types.Humanchain_Pool_Address {
		// Send true to HmPoolchange channel
		o.HmPoolChanged <- true
		return
	}

	if receiver != types.Humanchain_Pool_Address {
		return
	}

	// hash
	txHash := txs["tx.hash"][0]

	// amt
	amt := txs["transfer.amount"][0]
	amt = amt[:len(amt)-4]

	// convert uHMN to HMN
	famt, _ := strconv.ParseFloat(amt, 64)
	amount := fmt.Sprintf("%f", famt/1e9)

	_, voter := o.HumanChainBridge.GetVoterInfo()
	msg := types.NewMsgObservationVote(voter, txHash, types.CHAIN_HUMAN, sender, receiver, amount)
	o.ArrMsgObservationVote = append(o.ArrMsgObservationVote, msg)

	// Send true to HmPoolchange channel
	o.HmPoolChanged <- true
}
