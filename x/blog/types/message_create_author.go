package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAuthor{}

func NewMsgCreateAuthor(creator string, firstname string, lastname string) *MsgCreateAuthor {
	return &MsgCreateAuthor{
		Creator:   creator,
		Firstname: firstname,
		Lastname:  lastname,
	}
}

func (msg *MsgCreateAuthor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
