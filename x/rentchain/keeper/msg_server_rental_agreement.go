package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gopherine/rentchain/x/rentchain/types"
)

func (k msgServer) CreateRentalAgreement(goCtx context.Context, msg *types.MsgCreateRentalAgreement) (*types.MsgCreateRentalAgreementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetRentalAgreement(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var rentalAgreement = types.RentalAgreement{
		Creator:   msg.Creator,
		Index:     msg.Index,
		ItemId:    msg.ItemId,
		OwnerId:   msg.OwnerId,
		RenterId:  msg.RenterId,
		Price:     msg.Price,
		StartTime: msg.StartTime,
		Duration:  msg.Duration,
		IsActive:  msg.IsActive,
	}

	k.SetRentalAgreement(
		ctx,
		rentalAgreement,
	)
	return &types.MsgCreateRentalAgreementResponse{}, nil
}

func (k msgServer) UpdateRentalAgreement(goCtx context.Context, msg *types.MsgUpdateRentalAgreement) (*types.MsgUpdateRentalAgreementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetRentalAgreement(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var rentalAgreement = types.RentalAgreement{
		Creator:   msg.Creator,
		Index:     msg.Index,
		ItemId:    msg.ItemId,
		OwnerId:   msg.OwnerId,
		RenterId:  msg.RenterId,
		Price:     msg.Price,
		StartTime: msg.StartTime,
		Duration:  msg.Duration,
		IsActive:  msg.IsActive,
	}

	k.SetRentalAgreement(ctx, rentalAgreement)

	return &types.MsgUpdateRentalAgreementResponse{}, nil
}

func (k msgServer) DeleteRentalAgreement(goCtx context.Context, msg *types.MsgDeleteRentalAgreement) (*types.MsgDeleteRentalAgreementResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetRentalAgreement(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveRentalAgreement(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteRentalAgreementResponse{}, nil
}
