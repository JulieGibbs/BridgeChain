package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTranfserPoolcoin = "tranfser_poolcoin"

var _ sdk.Msg = &MsgTranfserPoolcoin{}

func NewMsgTranfserPoolcoin(creator string, addr string, amt string) *MsgTranfserPoolcoin {
	return &MsgTranfserPoolcoin{
		Creator: creator,
		Addr:    addr,
		Amt:     amt,
	}
}

func (msg *MsgTranfserPoolcoin) Route() string {
	return RouterKey
}

func (msg *MsgTranfserPoolcoin) Type() string {
	return TypeMsgTranfserPoolcoin
}

func (msg *MsgTranfserPoolcoin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTranfserPoolcoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTranfserPoolcoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
