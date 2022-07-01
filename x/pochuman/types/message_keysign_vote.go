package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgKeysignVote = "keysign_vote"

var _ sdk.Msg = &MsgKeysignVote{}

func NewMsgKeysignVote(creator string, txHash string, pubKey string) *MsgKeysignVote {
	return &MsgKeysignVote{
		Creator: creator,
		TxHash:  txHash,
		PubKey:  pubKey,
	}
}

func (msg *MsgKeysignVote) Route() string {
	return RouterKey
}

func (msg *MsgKeysignVote) Type() string {
	return TypeMsgKeysignVote
}

func (msg *MsgKeysignVote) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgKeysignVote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgKeysignVote) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
