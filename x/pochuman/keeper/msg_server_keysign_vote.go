package keeper

import (
	"context"
	"fmt"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) KeysignVote(goCtx context.Context, msg *types.MsgKeysignVote) (*types.MsgKeysignVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	//
	n := k.GetKeysignTxStoreCount(ctx)
	index := fmt.Sprintf("%d", n+1)

	// Create a new instance of ObserveVote
	transaction := types.KeysignVoteData{
		Index:  index,
		TxHash: msg.TxHash,
		PubKey: msg.PubKey,
		Voter:  msg.Creator,
		TxTime: ctx.BlockTime().UTC().String(),
	}

	// Add new keysign data
	k.SetKeysignVoteData(ctx, transaction)

	// Update Requested Vote Count
	k.SetKeysignTxStoreCount(ctx, n+1)

	return &types.MsgKeysignVoteResponse{Code: "200", Msg: "Operation succeed"}, nil
}
