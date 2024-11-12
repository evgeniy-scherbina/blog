package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdatePost{}

func NewMsgUpdatePost(creator string, title string, body string, id uint64) *MsgUpdatePost {
	return &MsgUpdatePost{
		Creator: creator,
		Title:   title,
		Body:    body,
		Id:      id,
	}
}

func (msg *MsgUpdatePost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
