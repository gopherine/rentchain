package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateRental{}

func NewMsgCreateRental(creator string, itemId string, ownerId string, renterId string, price string, startTime string, duration string) *MsgCreateRental {
	return &MsgCreateRental{
		Creator:   creator,
		ItemId:    itemId,
		OwnerId:   ownerId,
		RenterId:  renterId,
		Price:     price,
		StartTime: startTime,
		Duration:  duration,
	}
}

func (msg *MsgCreateRental) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
