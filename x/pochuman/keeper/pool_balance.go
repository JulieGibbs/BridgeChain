package keeper

import (
	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPoolBalance set a specific poolBalance in the store from its index
func (k Keeper) SetPoolBalance(ctx sdk.Context, poolBalance types.PoolBalance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolBalanceKeyPrefix))
	b := k.cdc.MustMarshal(&poolBalance)
	store.Set(types.PoolBalanceKey(
		poolBalance.Index,
	), b)
}

// GetPoolBalance returns a poolBalance from its index
func (k Keeper) GetPoolBalance(
	ctx sdk.Context,
	index string,

) (val types.PoolBalance, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolBalanceKeyPrefix))

	b := store.Get(types.PoolBalanceKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePoolBalance removes a poolBalance from the store
func (k Keeper) RemovePoolBalance(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolBalanceKeyPrefix))
	store.Delete(types.PoolBalanceKey(
		index,
	))
}

// GetAllPoolBalance returns all poolBalance
func (k Keeper) GetAllPoolBalance(ctx sdk.Context) (list []types.PoolBalance) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PoolBalanceKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PoolBalance
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
