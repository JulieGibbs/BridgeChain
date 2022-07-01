package keeper

import (
	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetTransactionData set a specific transactionData in the store from its index
func (k Keeper) SetTransactionData(ctx sdk.Context, transactionData types.TransactionData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionDataKeyPrefix))
	b := k.cdc.MustMarshal(&transactionData)
	store.Set(types.TransactionDataKey(
		transactionData.Index,
	), b)
}

// GetTransactionData returns a transactionData from its index
func (k Keeper) GetTransactionData(
	ctx sdk.Context,
	index string,

) (val types.TransactionData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionDataKeyPrefix))

	b := store.Get(types.TransactionDataKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTransactionData removes a transactionData from the store
func (k Keeper) RemoveTransactionData(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionDataKeyPrefix))
	store.Delete(types.TransactionDataKey(
		index,
	))
}

// GetAllTransactionData returns all transactionData
func (k Keeper) GetAllTransactionData(ctx sdk.Context) (list []types.TransactionData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionDataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TransactionData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
