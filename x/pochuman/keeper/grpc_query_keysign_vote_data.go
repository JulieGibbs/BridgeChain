package keeper

import (
	"context"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) KeysignVoteDataAll(c context.Context, req *types.QueryAllKeysignVoteDataRequest) (*types.QueryAllKeysignVoteDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var keysignVoteDatas []types.KeysignVoteData
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	keysignVoteDataStore := prefix.NewStore(store, types.KeyPrefix(types.KeysignVoteDataKeyPrefix))

	pageRes, err := query.Paginate(keysignVoteDataStore, req.Pagination, func(key []byte, value []byte) error {
		var keysignVoteData types.KeysignVoteData
		if err := k.cdc.Unmarshal(value, &keysignVoteData); err != nil {
			return err
		}

		keysignVoteDatas = append(keysignVoteDatas, keysignVoteData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllKeysignVoteDataResponse{KeysignVoteData: keysignVoteDatas, Pagination: pageRes}, nil
}

func (k Keeper) KeysignVoteData(c context.Context, req *types.QueryGetKeysignVoteDataRequest) (*types.QueryGetKeysignVoteDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetKeysignVoteData(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetKeysignVoteDataResponse{KeysignVoteData: val}, nil
}
