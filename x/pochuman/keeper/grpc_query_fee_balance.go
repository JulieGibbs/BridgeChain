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

func (k Keeper) FeeBalanceAll(c context.Context, req *types.QueryAllFeeBalanceRequest) (*types.QueryAllFeeBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var feeBalances []types.FeeBalance
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	feeBalanceStore := prefix.NewStore(store, types.KeyPrefix(types.FeeBalanceKeyPrefix))

	pageRes, err := query.Paginate(feeBalanceStore, req.Pagination, func(key []byte, value []byte) error {
		var feeBalance types.FeeBalance
		if err := k.cdc.Unmarshal(value, &feeBalance); err != nil {
			return err
		}

		feeBalances = append(feeBalances, feeBalance)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFeeBalanceResponse{FeeBalance: feeBalances, Pagination: pageRes}, nil
}

func (k Keeper) FeeBalance(c context.Context, req *types.QueryGetFeeBalanceRequest) (*types.QueryGetFeeBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetFeeBalance(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetFeeBalanceResponse{FeeBalance: val}, nil
}
