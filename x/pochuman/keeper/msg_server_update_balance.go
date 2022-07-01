package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateBalance(goCtx context.Context, msg *types.MsgUpdateBalance) (*types.MsgUpdateBalanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	index := "2"
	chainName := "Human"
	if msg.ChainName == types.CHAIN_ETHEREUM {
		index = "1"
		chainName = "Ethereum"
	}

	// if it is a wrong chain id return
	if msg.ChainName != types.CHAIN_ETHEREUM && msg.ChainName != types.CHAIN_HUMAN {
		return &types.MsgUpdateBalanceResponse{Code: "501", Msg: "Invalid amount parameter"}, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid chain")
	}

	amt, err := strconv.ParseFloat(msg.Balance, 64)
	if err != nil || amt <= 0 {
		return &types.MsgUpdateBalanceResponse{Code: "501", Msg: "Invalid amount parameter"}, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid coins (%s)", err)
	}

	// Minus Fee Amount
	fFee := float64(0.0)
	feeItem, bFoundF := k.GetFeeBalance(ctx, index)
	if bFoundF {
		fFee, _ = strconv.ParseFloat(feeItem.Balance, 64)
	}

	fBalance, _ := strconv.ParseFloat(msg.Balance, 64)
	fBalance -= fFee

	// Create a new instance of Bet
	balance := types.PoolBalance{
		Index:     index,
		ChainName: chainName,
		Balance:   fmt.Sprintf("%f", fBalance),
		Decimal:   msg.Decimal,
	}

	// Update Pool Balance
	k.SetPoolBalance(ctx, balance)

	return &types.MsgUpdateBalanceResponse{}, nil
}
