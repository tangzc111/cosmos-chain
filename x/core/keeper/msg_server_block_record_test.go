package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"cosmos-chain/x/core/keeper"
	"cosmos-chain/x/core/types"
)

func TestBlockRecordMsgServerCreate(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)
	creator, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateBlockRecord{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateBlockRecord(f.ctx, expected)
		require.NoError(t, err)
		rst, err := f.keeper.BlockRecord.Get(f.ctx, expected.Index)
		require.NoError(t, err)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestBlockRecordMsgServerUpdate(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	unauthorizedAddr, err := f.addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)

	expected := &types.MsgCreateBlockRecord{Creator: creator,
		Index: strconv.Itoa(0),
	}
	_, err = srv.CreateBlockRecord(f.ctx, expected)
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgUpdateBlockRecord
		err     error
	}{
		{
			desc: "invalid address",
			request: &types.MsgUpdateBlockRecord{Creator: "invalid",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			desc: "unauthorized",
			request: &types.MsgUpdateBlockRecord{Creator: unauthorizedAddr,
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "key not found",
			request: &types.MsgUpdateBlockRecord{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "completed",
			request: &types.MsgUpdateBlockRecord{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.UpdateBlockRecord(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, err := f.keeper.BlockRecord.Get(f.ctx, expected.Index)
				require.NoError(t, err)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestBlockRecordMsgServerDelete(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)

	unauthorizedAddr, err := f.addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)

	_, err = srv.CreateBlockRecord(f.ctx, &types.MsgCreateBlockRecord{Creator: creator,
		Index: strconv.Itoa(0),
	})
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgDeleteBlockRecord
		err     error
	}{
		{
			desc: "invalid address",
			request: &types.MsgDeleteBlockRecord{Creator: "invalid",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			desc: "unauthorized",
			request: &types.MsgDeleteBlockRecord{Creator: unauthorizedAddr,
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "key not found",
			request: &types.MsgDeleteBlockRecord{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "completed",
			request: &types.MsgDeleteBlockRecord{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.DeleteBlockRecord(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				found, err := f.keeper.BlockRecord.Has(f.ctx, tc.request.Index)
				require.NoError(t, err)
				require.False(t, found)
			}
		})
	}
}
