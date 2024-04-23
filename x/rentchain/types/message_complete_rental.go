package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCompleteRental{}

func NewMsgCompleteRental(creator string, itemId string) *MsgCompleteRental {
	return &MsgCompleteRental{
		Creator: creator,
		ItemId:  itemId,
	}
}

func (msg *MsgCompleteRental) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
