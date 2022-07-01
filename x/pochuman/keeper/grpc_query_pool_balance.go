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

func (k Keeper) PoolBalanceAll(c context.Context, req *types.QueryAllPoolBalanceRequest) (*types.QueryAllPoolBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var poolBalances []types.PoolBalance
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	poolBalanceStore := prefix.NewStore(store, types.KeyPrefix(types.PoolBalanceKeyPrefix))

	pageRes, err := query.Paginate(poolBalanceStore, req.Pagination, func(key []byte, value []byte) error {
		var poolBalance types.PoolBalance
		if err := k.cdc.Unmarshal(value, &poolBalance); err != nil {
			return err
		}

		poolBalances = append(poolBalances, poolBalance)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPoolBalanceResponse{PoolBalance: poolBalances, Pagination: pageRes}, nil
}

func (k Keeper) PoolBalance(c context.Context, req *types.QueryGetPoolBalanceRequest) (*types.QueryGetPoolBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPoolBalance(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPoolBalanceResponse{PoolBalance: val}, nil
}
