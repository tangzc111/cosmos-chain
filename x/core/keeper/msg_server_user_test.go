package keeper_test

import (
	"fmt"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"cosmos-chain/x/core/keeper"
	"cosmos-chain/x/core/types"
)

func TestUserMsgServerCreate(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	for i := 0; i < 5; i++ {
		addr, err := f.addressCodec.BytesToString([]byte(fmt.Sprintf("userAddr_%02d______________", i)))
		require.NoError(t, err)
		expected := &types.MsgCreateUser{
			Creator:     addr,
			Index:       addr,
			Address:     addr,
			Username:    fmt.Sprintf("user-%d", i),
			Description: "test user",
		}
		_, err = srv.CreateUser(f.ctx, expected)
		require.NoError(t, err)
		rst, err := f.keeper.User.Get(f.ctx, expected.Index)
		require.NoError(t, err)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestUserMsgServerUpdate(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("userAddr_creator___________"))
	require.NoError(t, err)

	unauthorizedAddr, err := f.addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)
	missingAddr, err := f.addressCodec.BytesToString([]byte("missingAddr______________"))
	require.NoError(t, err)

	expected := &types.MsgCreateUser{
		Creator:     creator,
		Index:       creator,
		Address:     creator,
		Username:    "original",
		Description: "owner",
	}
	_, err = srv.CreateUser(f.ctx, expected)
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgUpdateUser
		err     error
	}{
		{
			desc: "invalid address",
			request: &types.MsgUpdateUser{Creator: "invalid",
				Index: creator,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			desc: "unauthorized",
			request: &types.MsgUpdateUser{Creator: unauthorizedAddr,
				Index: creator,
			},
			err: types.ErrUnauthorized,
		},
		{
			desc: "key not found",
			request: &types.MsgUpdateUser{Creator: creator,
				Index: missingAddr,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "completed",
			request: &types.MsgUpdateUser{Creator: creator,
				Index: creator,
				Username: "updated",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.UpdateUser(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, err := f.keeper.User.Get(f.ctx, expected.Index)
				require.NoError(t, err)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestUserMsgServerDelete(t *testing.T) {
	f := initFixture(t)
	srv := keeper.NewMsgServerImpl(f.keeper)

	creator, err := f.addressCodec.BytesToString([]byte("userAddr_creator___________"))
	require.NoError(t, err)

	unauthorizedAddr, err := f.addressCodec.BytesToString([]byte("unauthorizedAddr___________"))
	require.NoError(t, err)
	missingAddr, err := f.addressCodec.BytesToString([]byte("missingAddr______________"))
	require.NoError(t, err)

	_, err = srv.CreateUser(f.ctx, &types.MsgCreateUser{
		Creator:     creator,
		Index:       creator,
		Address:     creator,
		Username:    "original",
		Description: "owner",
	})
	require.NoError(t, err)

	tests := []struct {
		desc    string
		request *types.MsgDeleteUser
		err     error
	}{
		{
			desc: "invalid address",
			request: &types.MsgDeleteUser{Creator: "invalid",
				Index: creator,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			desc: "unauthorized",
			request: &types.MsgDeleteUser{Creator: unauthorizedAddr,
				Index: creator,
			},
			err: types.ErrUnauthorized,
		},
		{
			desc: "key not found",
			request: &types.MsgDeleteUser{Creator: creator,
				Index: missingAddr,
			},
			err: sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "completed",
			request: &types.MsgDeleteUser{Creator: creator,
				Index: creator,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, err = srv.DeleteUser(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				found, err := f.keeper.User.Has(f.ctx, tc.request.Index)
				require.NoError(t, err)
				require.False(t, found)
			}
		})
	}
}
