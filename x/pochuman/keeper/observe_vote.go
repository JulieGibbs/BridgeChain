package keeper

import (
	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetObserveVote set a specific observeVote in the store from its index
func (k Keeper) SetObserveVote(ctx sdk.Context, observeVote types.ObserveVote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObserveVoteKeyPrefix))
	b := k.cdc.MustMarshal(&observeVote)
	store.Set(types.ObserveVoteKey(
		observeVote.Index,
	), b)
}

// GetObserveVote returns a observeVote from its index
func (k Keeper) GetObserveVote(
	ctx sdk.Context,
	index string,

) (val types.ObserveVote, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObserveVoteKeyPrefix))

	b := store.Get(types.ObserveVoteKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveObserveVote removes a observeVote from the store
func (k Keeper) RemoveObserveVote(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObserveVoteKeyPrefix))
	store.Delete(types.ObserveVoteKey(
		index,
	))
}

// GetAllObserveVote returns all observeVote
func (k Keeper) GetAllObserveVote(ctx sdk.Context) (list []types.ObserveVote) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ObserveVoteKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.ObserveVote
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
