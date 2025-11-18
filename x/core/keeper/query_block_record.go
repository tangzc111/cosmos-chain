package keeper

import (
	"context"
	"errors"

	"tokenchain/x/core/types"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListBlockRecord(ctx context.Context, req *types.QueryAllBlockRecordRequest) (*types.QueryAllBlockRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	blockRecords, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.BlockRecord,
		req.Pagination,
		func(_ string, value types.BlockRecord) (types.BlockRecord, error) {
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBlockRecordResponse{BlockRecord: blockRecords, Pagination: pageRes}, nil
}

func (q queryServer) GetBlockRecord(ctx context.Context, req *types.QueryGetBlockRecordRequest) (*types.QueryGetBlockRecordResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.BlockRecord.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetBlockRecordResponse{BlockRecord: val}, nil
}
