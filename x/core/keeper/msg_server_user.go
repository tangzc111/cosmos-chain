package keeper

import (
	"context"
	"errors"

	"cosmos-chain/x/core/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateUser(ctx context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	if _, err := k.mustValidAddress(msg.Creator); err != nil {
		return nil, err
	}
	if _, err := k.mustValidAddress(msg.Address); err != nil {
		return nil, err
	}
	if msg.Index != "" && msg.Index != msg.Address {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index must match address")
	}
	if msg.Username == "" {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "username cannot be empty")
	}

	addrKey := k.normalizeUserKey(msg.Address)
	if exists, err := k.User.Has(ctx, addrKey); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if exists {
		return nil, types.ErrUserExists
	}

	var user = types.User{
		Creator:     msg.Creator,
		Index:       addrKey,
		Address:     msg.Address,
		Username:    msg.Username,
		Description: msg.Description,
	}

	if err := k.User.Set(ctx, addrKey, user); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateUserResponse{}, nil
}

func (k msgServer) UpdateUser(ctx context.Context, msg *types.MsgUpdateUser) (*types.MsgUpdateUserResponse, error) {
	if _, err := k.mustValidAddress(msg.Creator); err != nil {
		return nil, err
	}

	key := k.normalizeUserKey(msg.Index)

	// Check if the value exists
	val, err := k.User.Get(ctx, key)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "incorrect owner")
	}

	if msg.Address != "" && msg.Address != val.Address {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "address cannot be updated")
	}

	var user = types.User{
		Creator: msg.Creator,
		Index:   val.Index,
		Address: val.Address,
	}

	if msg.Username != "" {
		user.Username = msg.Username
	} else {
		user.Username = val.Username
	}

	if msg.Description != "" {
		user.Description = msg.Description
	} else {
		user.Description = val.Description
	}

	if err := k.User.Set(ctx, user.Index, user); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update user")
	}

	return &types.MsgUpdateUserResponse{}, nil
}

func (k msgServer) DeleteUser(ctx context.Context, msg *types.MsgDeleteUser) (*types.MsgDeleteUserResponse, error) {
	if _, err := k.mustValidAddress(msg.Creator); err != nil {
		return nil, err
	}

	// Check if the value exists
	val, err := k.User.Get(ctx, k.normalizeUserKey(msg.Index))
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "incorrect owner")
	}

	if err := k.User.Remove(ctx, val.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove user")
	}

	return &types.MsgDeleteUserResponse{}, nil
}
