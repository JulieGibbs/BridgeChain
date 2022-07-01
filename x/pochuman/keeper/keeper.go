package keeper

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Read requested transaction data count
func (k Keeper) GetReqTransactionStoreCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ReqTransactionCountStoreKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// Write requested transaction data count
func (k Keeper) SetReqTransactionStoreCount(ctx sdk.Context, count uint64) bool {
	store := ctx.KVStore(k.storeKey)
	bz := sdk.Uint64ToBigEndian(count)
	store.Set(types.ReqTransactionCountStoreKey, bz)

	return true
}

// Read requested observe data count
func (k Keeper) GetObserveTxStoreCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ObserveTxCountStoreKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// Write requested observe data count
func (k Keeper) SetObserveTxStoreCount(ctx sdk.Context, count uint64) bool {
	store := ctx.KVStore(k.storeKey)
	bz := sdk.Uint64ToBigEndian(count)
	store.Set(types.ObserveTxCountStoreKey, bz)

	return true
}

// Read requested observe data count
func (k Keeper) GetKeysignTxStoreCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeysignTxCountStoreKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// Write requested observe data count
func (k Keeper) SetKeysignTxStoreCount(ctx sdk.Context, count uint64) bool {
	store := ctx.KVStore(k.storeKey)
	bz := sdk.Uint64ToBigEndian(count)
	store.Set(types.KeysignTxCountStoreKey, bz)

	return true
}

// Update transaction data from the eth transaction history
func (k Keeper) KeysignTxData(ctx sdk.Context, ksVotes []types.KeysignVoteData, d types.TransactionData) bool {
	nVoteCnt := 0
	// looping
	for _, data := range ksVotes {
		if data.TxHash != d.ConfirmedBlockHash {
			continue
		}

		nVoteCnt++
	}

	// Update Block Hash and Status
	if nVoteCnt >= types.VALID_CNT_SIGNER {
		d.Status = types.PAY_KEYSIGNED
	}

	// Update store
	k.SetTransactionData(ctx, d)

	return false
}

// Each time, we check observation votes and update tx requst appropriately.
// Later we will introduce node slashing logic here
func (k Keeper) UpdateTxRequestByObservationVote(ctx sdk.Context) bool {
	// Get all observation votes
	obsVotes := k.GetAllObserveVote(ctx)

	availableObs := make([]types.ObserveVote, 0)
	t := ctx.BlockTime().UTC()

	// Get available observation votes list
	// Reverse looping
	for i := len(obsVotes) - 1; i >= 0; i-- {
		ov := obsVotes[i]

		layout := "2006-01-02 15:04:05 -0700 MST"
		ot, _ := time.Parse(layout, ov.Txtime)
		ds := t.Sub(ot).Seconds()

		// Timed out, break
		if ds > types.VOTE_DIFF || ov.Used == "true" {
			continue
		}

		// append
		availableObs = append(availableObs, ov)
	}

	// Get all transaction data
	txsData := k.GetAllTransactionData(ctx)

	// Update transaction data by observation vote
	for i := len(txsData) - 1; i >= 0; i-- {
		txD := txsData[i]

		// Skip confirmed transaction
		if txD.ConfirmedBlockHash != "" {
			continue
		}

		k.UpdateTransactionRequestData(ctx, availableObs, txD)
	}

	return true
}

// Sign the transaction so that we can trigger external transaction
// Later we will introduce node slashing logic here
func (k Keeper) KeysignTxRequest(ctx sdk.Context) bool {
	// Get all keysign votes
	kSignVotes := k.GetAllKeysignVoteData(ctx)

	availableObs := make([]types.KeysignVoteData, 0)
	t := ctx.BlockTime().UTC()

	// Get available keysign votes list
	// Reverse looping
	for i := len(kSignVotes) - 1; i >= 0; i-- {
		kv := kSignVotes[i]

		layout := "2006-01-02 15:04:05 -0700 MST"
		ot, _ := time.Parse(layout, kv.TxTime)
		ds := t.Sub(ot).Seconds()

		// Timed out, break
		if ds > types.VOTE_DIFF {
			continue
		}

		// append
		availableObs = append(availableObs, kv)
	}

	// Get all transaction data
	txsData := k.GetAllTransactionData(ctx)

	// Update transaction data by observation vote
	for i := len(txsData) - 1; i >= 0; i-- {
		txD := txsData[i]

		// Skip unconfirmed transaction
		if txD.Status != types.PAY_CONFIRMED {
			continue
		}

		// Sign the transaction
		k.KeysignTxData(ctx, availableObs, txD)
	}

	return true
}

// Update transaction data from the eth transaction history
func (k Keeper) UpdateTransactionRequestData(ctx sdk.Context, obsVotes []types.ObserveVote, d types.TransactionData) bool {
	nVoteCnt := 0
	if len(obsVotes) < 1 {
		return false
	}

	index := 0

	ids := make([]int, 0)
	// looping
	for id, data := range obsVotes {
		amt1, _ := strconv.ParseFloat(d.Amount, 64)
		amt2, _ := strconv.ParseFloat(data.Amount, 64)

		// Source chain compare
		if d.OriginChain != data.ChainId {
			continue
		}

		// if original address or amount token dismatches, continue
		if d.OriginAddress != data.From || amt1 != amt2 {
			continue
		}

		i, _ := strconv.ParseUint(data.Index, 10, 64)
		if int(i) > index {
			d.ConfirmedBlockHash = data.Txhash
			index = int(i)
		}

		ids = append(ids, id)
		nVoteCnt++
	}

	// Update Block Hash and Status
	if nVoteCnt >= types.VALID_CNT_SIGNER {
		d.Status = types.PAY_CONFIRMED
	} else {
		d.Status = types.PAY_UNAVAILABLE
	}

	// Update store
	k.SetTransactionData(ctx, d)

	// Update used status
	if d.Status == types.PAY_CONFIRMED {
		for _, i := range ids {
			obsVotes[i].Used = "true"
			k.SetObserveVote(ctx, obsVotes[i])
		}

		// After we confirm payment on external blockchain, we deduct fee
		index := "2"
		chainName := "Human"
		if d.OriginChain == types.CHAIN_ETHEREUM {
			index = "1"
			chainName = "Ethereum"
		}

		feeItem, bFound := k.GetFeeBalance(ctx, index)
		if !bFound {
			// Create a new instance of Bet
			balance := types.FeeBalance{
				Index:     index,
				ChainName: chainName,
				Balance:   d.Fee,
			}

			// Update Fee Balance
			k.SetFeeBalance(ctx, balance)
		} else {
			fPrev, _ := strconv.ParseFloat(feeItem.Balance, 64)
			fFee, _ := strconv.ParseFloat(d.Fee, 64)

			fFee += fPrev
			feeItem.Balance = fmt.Sprintf("%f", fFee)
			k.SetFeeBalance(ctx, feeItem)
		}
	}

	return false
}
