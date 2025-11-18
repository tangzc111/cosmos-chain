package keeper_test

import (
	"fmt"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"tokenchain/x/core/keeper"
	"tokenchain/x/core/types"
)

func TestMinerMsgServerCreate(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	for i := 0; i < 5; i++ {
		addr, err := f.addressCodec.BytesToString([]byte(fmt.Sprintf("minerAddr_%02d_____________", i)))
		require.NoError(t, err)
		_, err = srv.CreateUser(f.ctx, &types.MsgCreateUser{
			Creator:     addr,
			Index:       addr,
			Address:     addr,
			Username:    fmt.Sprintf("miner-%d", i),
			Description: "miner",
		})
		require.NoError(t, err)

		expected := &types.MsgCreateMiner{
			Creator:     addr,
			Index:       addr,
			Address:     addr,
			Power:       "1",
			Description: "miner",
		}
		_, err = srv.CreateMiner(f.ctx, expected)
		require.NoError(t, err)
		rst, err := f.keeper.Miner.Get(f.ctx, expected.Index)
		require.NoError(t, err)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestMinerMsgServerUpdate(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("minerAddr_creator__________"))
	require.NoError(t, err)

	unauthorizedAddr, err := f.addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)

	missingMiner, err := f.addressCodec.BytesToString([]byte("missingMiner______________"))
	require.NoError(t, err)

	_, err = srv.CreateUser(f.ctx, &types.MsgCreateUser{
		Creator:     creator,
		Index:       creator,
		Address:     creator,
		Username:    "miner-owner",
		Description: "miner",
	})
	require.NoError(t, err)

	expected := &types.MsgCreateMiner{
		Creator:     creator,
		Index:       creator,
		Address:     creator,
		Power:       "1",
		Description: "miner",
	}
	_, err = srv.CreateMiner(f.ctx, expected)
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgUpdateMiner
		err     error
	}{
		{
			desc: "invalid address",
			request: &types.MsgUpdateMiner{Creator: "invalid",
				Index: creator,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			desc: "unauthorized",
			request: &types.MsgUpdateMiner{Creator: unauthorizedAddr,
				Index: creator,
			},
			err: types.ErrUnauthorized,
		},
		{
			desc: "key not found",
			request: &types.MsgUpdateMiner{Creator: creator,
				Index: missingMiner,
			},
			err: types.ErrMinerNotFound,
		},
		{
			desc: "completed",
			request: &types.MsgUpdateMiner{Creator: creator,
				Index: creator,
				Power: "5",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.UpdateMiner(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, err := f.keeper.Miner.Get(f.ctx, expected.Index)
				require.NoError(t, err)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestMinerMsgServerDelete(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("minerAddr_creator__________"))
	require.NoError(t, err)

	unauthorizedAddr, err := f.addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)

	missingMiner, err := f.addressCodec.BytesToString([]byte("missingMiner______________"))
	require.NoError(t, err)

	_, err = srv.CreateUser(f.ctx, &types.MsgCreateUser{
		Creator:     creator,
		Index:       creator,
		Address:     creator,
		Username:    "miner-owner",
		Description: "miner",
	})
	require.NoError(t, err)

	_, err = srv.CreateMiner(f.ctx, &types.MsgCreateMiner{
		Creator:     creator,
		Index:       creator,
		Address:     creator,
		Power:       "1",
		Description: "miner",
	})
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgDeleteMiner
		err     error
	}{
		{
			desc: "invalid address",
			request: &types.MsgDeleteMiner{Creator: "invalid",
				Index: creator,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			desc: "unauthorized",
			request: &types.MsgDeleteMiner{Creator: unauthorizedAddr,
				Index: creator,
			},
			err: types.ErrUnauthorized,
		},
		{
			desc: "key not found",
			request: &types.MsgDeleteMiner{Creator: creator,
				Index: missingMiner,
			},
			err: types.ErrMinerNotFound,
		},
		{
			desc: "completed",
			request: &types.MsgDeleteMiner{Creator: creator,
				Index: creator,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.DeleteMiner(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				found, err := f.keeper.Miner.Has(f.ctx, tc.request.Index)
				require.NoError(t, err)
				require.False(t, found)
			}
		})
	}
}
