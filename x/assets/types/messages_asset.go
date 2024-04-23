package types

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/golang/protobuf/ptypes/any"
)

var _ sdk.Msg = &MsgCreateAsset{}

func NewMsgCreateAsset(creator string, owner string, name string, description string, details *any.Any, pricePerUnit string, unit string, tags []string) *MsgCreateAsset {
	anyDetails, _ := types.NewAnyWithValue(details)

	return &MsgCreateAsset{
		Creator:      creator,
		Owner:        owner,
		Name:         name,
		Description:  description,
		Details:      anyDetails,
		PricePerUnit: pricePerUnit,
		Unit:         unit,
		Tags:         tags,
	}
}

func (msg *MsgCreateAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAsset{}

func NewMsgUpdateAsset(creator string, id uint64, owner string, name string, description string, details *any.Any, pricePerUnit string, unit string, tags []string) *MsgUpdateAsset {
	anyDetails, _ := types.NewAnyWithValue(details)

	return &MsgUpdateAsset{
		Id:           id,
		Creator:      creator,
		Name:         name,
		Description:  description,
		Details:      anyDetails,
		PricePerUnit: pricePerUnit,
		Unit:         unit,
		Tags:         tags,
	}
}

func (msg *MsgUpdateAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAsset{}

func NewMsgDeleteAsset(creator string, id uint64) *MsgDeleteAsset {
	return &MsgDeleteAsset{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
