package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveTransaction = "approve_transaction"

var _ sdk.Msg = &MsgApproveTransaction{}

func NewMsgApproveTransaction(creator string, txHash string, success string, signedKey string) *MsgApproveTransaction {
	return &MsgApproveTransaction{
		Creator:   creator,
		TxHash:    txHash,
		Success:   success,
		SignedKey: signedKey,
	}
}

func (msg *MsgApproveTransaction) Route() string {
	return RouterKey
}

func (msg *MsgApproveTransaction) Type() string {
	return TypeMsgApproveTransaction
}

func (msg *MsgApproveTransaction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
