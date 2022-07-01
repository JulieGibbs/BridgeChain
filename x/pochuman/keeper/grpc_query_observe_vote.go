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

func (k Keeper) ObserveVoteAll(c context.Context, req *types.QueryAllObserveVoteRequest) (*types.QueryAllObserveVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var observeVotes []types.ObserveVote
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	observeVoteStore := prefix.NewStore(store, types.KeyPrefix(types.ObserveVoteKeyPrefix))

	pageRes, err := query.Paginate(observeVoteStore, req.Pagination, func(key []byte, value []byte) error {
		var observeVote types.ObserveVote
		if err := k.cdc.Unmarshal(value, &observeVote); err != nil {
			return err
		}

		observeVotes = append(observeVotes, observeVote)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllObserveVoteResponse{ObserveVote: observeVotes, Pagination: pageRes}, nil
}

func (k Keeper) ObserveVote(c context.Context, req *types.QueryGetObserveVoteRequest) (*types.QueryGetObserveVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetObserveVote(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetObserveVoteResponse{ObserveVote: val}, nil
}
