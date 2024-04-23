package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateRentalAgreement{}

func NewMsgCreateRentalAgreement(
	creator string,
	index string,
	itemId string,
	ownerId string,
	renterId string,
	price string,
	startTime string,
	duration string,
	isActive bool,

) *MsgCreateRentalAgreement {
	return &MsgCreateRentalAgreement{
		Creator:   creator,
		Index:     index,
		ItemId:    itemId,
		OwnerId:   ownerId,
		RenterId:  renterId,
		Price:     price,
		StartTime: startTime,
		Duration:  duration,
		IsActive:  isActive,
	}
}

func (msg *MsgCreateRentalAgreement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRentalAgreement{}

func NewMsgUpdateRentalAgreement(
	creator string,
	index string,
	itemId string,
	ownerId string,
	renterId string,
	price string,
	startTime string,
	duration string,
	isActive bool,

) *MsgUpdateRentalAgreement {
	return &MsgUpdateRentalAgreement{
		Creator:   creator,
		Index:     index,
		ItemId:    itemId,
		OwnerId:   ownerId,
		RenterId:  renterId,
		Price:     price,
		StartTime: startTime,
		Duration:  duration,
		IsActive:  isActive,
	}
}

func (msg *MsgUpdateRentalAgreement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteRentalAgreement{}

func NewMsgDeleteRentalAgreement(
	creator string,
	index string,

) *MsgDeleteRentalAgreement {
	return &MsgDeleteRentalAgreement{
		Creator: creator,
		Index:   index,
	}
}

func (msg *MsgDeleteRentalAgreement) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
