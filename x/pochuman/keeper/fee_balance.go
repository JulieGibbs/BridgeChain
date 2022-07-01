package keeper

import (
	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetFeeBalance set a specific feeBalance in the store from its index
func (k Keeper) SetFeeBalance(ctx sdk.Context, feeBalance types.FeeBalance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeBalanceKeyPrefix))
	b := k.cdc.MustMarshal(&feeBalance)
	store.Set(types.FeeBalanceKey(
		feeBalance.Index,
	), b)
}

// GetFeeBalance returns a feeBalance from its index
func (k Keeper) GetFeeBalance(
	ctx sdk.Context,
	index string,

) (val types.FeeBalance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeBalanceKeyPrefix))

	b := store.Get(types.FeeBalanceKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFeeBalance removes a feeBalance from the store
func (k Keeper) RemoveFeeBalance(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeBalanceKeyPrefix))
	store.Delete(types.FeeBalanceKey(
		index,
	))
}

// GetAllFeeBalance returns all feeBalance
func (k Keeper) GetAllFeeBalance(ctx sdk.Context) (list []types.FeeBalance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FeeBalanceKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FeeBalance
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
