package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestTransaction = "request_transaction"

var _ sdk.Msg = &MsgRequestTransaction{}

func NewMsgRequestTransaction(creator string, originChain string, originAddress string, targetChain string, targetAddress string, amount string, fee string) *MsgRequestTransaction {
	return &MsgRequestTransaction{
		Creator:       creator,
		OriginChain:   originChain,
		OriginAddress: originAddress,
		TargetChain:   targetChain,
		TargetAddress: targetAddress,
		Amount:        amount,
		Fee:           fee,
	}
}

func (msg *MsgRequestTransaction) Route() string {
	return RouterKey
}

func (msg *MsgRequestTransaction) Type() string {
	return TypeMsgRequestTransaction
}

func (msg *MsgRequestTransaction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
