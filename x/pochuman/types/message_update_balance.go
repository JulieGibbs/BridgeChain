package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateBalance = "update_balance"

var _ sdk.Msg = &MsgUpdateBalance{}

func NewMsgUpdateBalance(creator string, chainName string, balance string, decimal string) *MsgUpdateBalance {
	return &MsgUpdateBalance{
		Creator:   creator,
		ChainName: chainName,
		Balance:   balance,
		Decimal:   decimal,
	}
}

func (msg *MsgUpdateBalance) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBalance) Type() string {
	return TypeMsgUpdateBalance
}

func (msg *MsgUpdateBalance) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBalance) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBalance) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
