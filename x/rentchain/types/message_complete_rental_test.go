package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gopherine/rentchain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgCompleteRental_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCompleteRental
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCompleteRental{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCompleteRental{
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
