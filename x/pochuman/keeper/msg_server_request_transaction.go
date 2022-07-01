package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RequestTransaction(goCtx context.Context, msg *types.MsgRequestTransaction) (*types.MsgRequestTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	amt, err := strconv.ParseFloat(msg.Amount, 64)
	if err != nil || amt <= 0 {
		return &types.MsgRequestTransactionResponse{Code: "501", Msg: "Invalid amount parameter"}, nil
	}

	if msg.OriginChain != types.CHAIN_ETHEREUM && msg.OriginChain != types.CHAIN_HUMAN {
		return &types.MsgRequestTransactionResponse{Code: "501", Msg: "Invalid chain parameter"}, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid chain")
	}

	if msg.TargetChain != types.CHAIN_ETHEREUM && msg.TargetChain != types.CHAIN_HUMAN {
		return &types.MsgRequestTransactionResponse{Code: "501", Msg: "Invalid chain parameter"}, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid chain")
	}

	// check destination chain
	if msg.TargetChain == types.CHAIN_HUMAN {
		balance, er := k.GetPoolBalance(ctx, "2")
		if !er {
			return &types.MsgRequestTransactionResponse{Code: "501", Msg: "Insufficient balance"}, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Solana pool doens't have enough balance")
		}

		b, e := strconv.ParseFloat(balance.Balance, 64)
		if e != nil {
			return &types.MsgRequestTransactionResponse{Code: "501", Msg: "Insufficient balance"}, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Solana pool doens't have enough balance")
		}

		// if pool balance is lower than transfer amount
		if b < amt {
			return &types.MsgRequestTransactionResponse{Code: "501", Msg: "Insufficient balance"}, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Solana pool doens't have enough balance")
		}
	}

	// check dest chain
	if msg.TargetChain == types.CHAIN_ETHEREUM {
		balance, er := k.GetPoolBalance(ctx, "1")
		if !er {
			return &types.MsgRequestTransactionResponse{Code: "501", Msg: "Insufficient balance"}, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Ethereum pool doens't have enough balance")
		}

		b, e := strconv.ParseFloat(balance.Balance, 64)
		if e != nil {
			return &types.MsgRequestTransactionResponse{Code: "501", Msg: "Insufficient balance"}, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Ethereum pool doens't have enough balance")
		}

		// if pool balance is lower than transfer amount
		if b < amt {
			return &types.MsgRequestTransactionResponse{Code: "501", Msg: "Insufficient balance"}, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Ethereum pool doens't have enough balance")
		}
	}

	// Convert player address strings to sdk.AccAddress
	payer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// take service fee from the payer
	fee := sdk.Coins{sdk.NewInt64Coin("uHMN", 1e9)}
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, payer, types.ModuleName, fee)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid coins (%s)", err)
	}

	n := k.GetReqTransactionStoreCount(ctx)
	index := fmt.Sprintf("%d", n+1)

	// Create a new instance of Bet
	transaction := types.TransactionData{
		Index:              index,
		OriginChain:        msg.OriginChain,
		TargetChain:        msg.TargetChain,
		Amount:             msg.Amount,
		Time:               ctx.BlockTime().UTC().String(),
		Creator:            msg.Creator,
		OriginAddress:      msg.OriginAddress,
		TargetAddress:      msg.TargetAddress,
		ConfirmedBlockHash: "",
		Status:             types.PAY_AVAILABLE,
		SignedKey:          "",
		Fee:                msg.Fee,
	}

	// Add new transaction data
	k.SetTransactionData(ctx, transaction)

	// Update Requested Transaction Count
	k.SetReqTransactionStoreCount(ctx, n+1)

	return &types.MsgRequestTransactionResponse{Code: "200", Msg: "Operation succeed"}, nil
}
