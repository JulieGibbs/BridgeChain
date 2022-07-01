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

func (k Keeper) TransactionDataAll(c context.Context, req *types.QueryAllTransactionDataRequest) (*types.QueryAllTransactionDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var transactionDatas []types.TransactionData
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	transactionDataStore := prefix.NewStore(store, types.KeyPrefix(types.TransactionDataKeyPrefix))

	pageRes, err := query.Paginate(transactionDataStore, req.Pagination, func(key []byte, value []byte) error {
		var transactionData types.TransactionData
		if err := k.cdc.Unmarshal(value, &transactionData); err != nil {
			return err
		}

		transactionDatas = append(transactionDatas, transactionData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTransactionDataResponse{TransactionData: transactionDatas, Pagination: pageRes}, nil
}

func (k Keeper) TransactionData(c context.Context, req *types.QueryGetTransactionDataRequest) (*types.QueryGetTransactionDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetTransactionData(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTransactionDataResponse{TransactionData: val}, nil
}
