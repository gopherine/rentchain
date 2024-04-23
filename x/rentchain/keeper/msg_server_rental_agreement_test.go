package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/gopherine/rentchain/testutil/keeper"
	"github.com/gopherine/rentchain/x/rentchain/keeper"
	"github.com/gopherine/rentchain/x/rentchain/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRentalAgreementMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.RentchainKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateRentalAgreement{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateRentalAgreement(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetRentalAgreement(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestRentalAgreementMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateRentalAgreement
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateRentalAgreement{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateRentalAgreement{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateRentalAgreement{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RentchainKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateRentalAgreement{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateRentalAgreement(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateRentalAgreement(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetRentalAgreement(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestRentalAgreementMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteRentalAgreement
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteRentalAgreement{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteRentalAgreement{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteRentalAgreement{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.RentchainKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateRentalAgreement(ctx, &types.MsgCreateRentalAgreement{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteRentalAgreement(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetRentalAgreement(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
