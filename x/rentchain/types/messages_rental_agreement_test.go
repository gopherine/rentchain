package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gopherine/rentchain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateRentalAgreement_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateRentalAgreement
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateRentalAgreement{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateRentalAgreement{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateRentalAgreement_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateRentalAgreement
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateRentalAgreement{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateRentalAgreement{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteRentalAgreement_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteRentalAgreement
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteRentalAgreement{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteRentalAgreement{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
