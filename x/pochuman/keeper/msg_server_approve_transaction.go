package keeper

import (
	"context"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ApproveTransaction(goCtx context.Context, msg *types.MsgApproveTransaction) (*types.MsgApproveTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get all transactions to loop
	allData := k.GetAllTransactionData(ctx)

	// Loop transactions and process
	for i := len(allData) - 1; i >= 0; i-- {
		d := allData[i]

		// if it's not confirmed one, skip it
		if d.ConfirmedBlockHash != msg.TxHash {
			continue
		}

		// Updating the data state to paid
		d.Status = msg.Success
		d.SignedKey = msg.SignedKey
		k.SetTransactionData(ctx, d)

		// Return result with success
		return &types.MsgApproveTransactionResponse{Code: "200", Msg: "Operation succeed."}, nil
	}

	return &types.MsgApproveTransactionResponse{Code: "301", Msg: "No matched transactions found."}, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Internal error occurs while transferring the token.")
}
