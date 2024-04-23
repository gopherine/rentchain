package keeper_test

import (
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/gopherine/rentchain/testutil/keeper"
	"github.com/gopherine/rentchain/testutil/nullify"
	"github.com/gopherine/rentchain/x/rentchain/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestRentalAgreementQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.RentchainKeeper(t)
	msgs := createNRentalAgreement(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetRentalAgreementRequest
		response *types.QueryGetRentalAgreementResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetRentalAgreementRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetRentalAgreementResponse{RentalAgreement: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetRentalAgreementRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetRentalAgreementResponse{RentalAgreement: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetRentalAgreementRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.RentalAgreement(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestRentalAgreementQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.RentchainKeeper(t)
	msgs := createNRentalAgreement(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllRentalAgreementRequest {
		return &types.QueryAllRentalAgreementRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RentalAgreementAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RentalAgreement), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RentalAgreement),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.RentalAgreementAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.RentalAgreement), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.RentalAgreement),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.RentalAgreementAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.RentalAgreement),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.RentalAgreementAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
