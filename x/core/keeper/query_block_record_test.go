package keeper_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"tokenchain/x/core/keeper"
	"tokenchain/x/core/types"
)

func createNBlockRecord(keeper keeper.Keeper, ctx context.Context, n int) []types.BlockRecord {
	items := make([]types.BlockRecord, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Height = strconv.Itoa(i)
		items[i].Hash = strconv.Itoa(i)
		items[i].Proposer = strconv.Itoa(i)
		items[i].Time = strconv.Itoa(i)
		_ = keeper.BlockRecord.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestBlockRecordQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNBlockRecord(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetBlockRecordRequest
		response *types.QueryGetBlockRecordResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetBlockRecordRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetBlockRecordResponse{BlockRecord: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetBlockRecordRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetBlockRecordResponse{BlockRecord: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetBlockRecordRequest{
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
			response, err := qs.GetBlockRecord(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestBlockRecordQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNBlockRecord(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllBlockRecordRequest {
		return &types.QueryAllBlockRecordRequest{
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
			resp, err := qs.ListBlockRecord(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.BlockRecord), step)
			require.Subset(t, msgs, resp.BlockRecord)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListBlockRecord(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.BlockRecord), step)
			require.Subset(t, msgs, resp.BlockRecord)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListBlockRecord(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.BlockRecord)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListBlockRecord(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
