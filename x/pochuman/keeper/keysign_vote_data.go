package keeper

import (
	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetKeysignVoteData set a specific keysignVoteData in the store from its index
func (k Keeper) SetKeysignVoteData(ctx sdk.Context, keysignVoteData types.KeysignVoteData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeysignVoteDataKeyPrefix))
	b := k.cdc.MustMarshal(&keysignVoteData)
	store.Set(types.KeysignVoteDataKey(
		keysignVoteData.Index,
	), b)
}

// GetKeysignVoteData returns a keysignVoteData from its index
func (k Keeper) GetKeysignVoteData(
	ctx sdk.Context,
	index string,

) (val types.KeysignVoteData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeysignVoteDataKeyPrefix))

	b := store.Get(types.KeysignVoteDataKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveKeysignVoteData removes a keysignVoteData from the store
func (k Keeper) RemoveKeysignVoteData(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeysignVoteDataKeyPrefix))
	store.Delete(types.KeysignVoteDataKey(
		index,
	))
}

// GetAllKeysignVoteData returns all keysignVoteData
func (k Keeper) GetAllKeysignVoteData(ctx sdk.Context) (list []types.KeysignVoteData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.KeysignVoteDataKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.KeysignVoteData
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
