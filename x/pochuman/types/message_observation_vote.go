package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgObservationVote = "observation_vote"

var _ sdk.Msg = &MsgObservationVote{}

func NewMsgObservationVote(creator string, txHash string, chainName string, from string, to string, amount string) *MsgObservationVote {
	return &MsgObservationVote{
		Creator:   creator,
		TxHash:    txHash,
		ChainName: chainName,
		From:      from,
		To:        to,
		Amount:    amount,
	}
}

func (msg *MsgObservationVote) Route() string {
	return RouterKey
}

func (msg *MsgObservationVote) Type() string {
	return TypeMsgObservationVote
}

func (msg *MsgObservationVote) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgObservationVote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgObservationVote) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
