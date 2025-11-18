package keeper

import (
	"context"
	"errors"
	"fmt"

	"tokenchain/x/core/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateUser(ctx context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.User.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var user = types.User{
		Creator:     msg.Creator,
		Index:       msg.Index,
		Address:     msg.Address,
		Username:    msg.Username,
		Description: msg.Description,
	}

	if err := k.User.Set(ctx, user.Index, user); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateUserResponse{}, nil
}

func (k msgServer) UpdateUser(ctx context.Context, msg *types.MsgUpdateUser) (*types.MsgUpdateUserResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.User.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var user = types.User{
		Creator:     msg.Creator,
		Index:       msg.Index,
		Address:     msg.Address,
		Username:    msg.Username,
		Description: msg.Description,
	}

	if err := k.User.Set(ctx, user.Index, user); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update user")
	}

	return &types.MsgUpdateUserResponse{}, nil
}

func (k msgServer) DeleteUser(ctx context.Context, msg *types.MsgDeleteUser) (*types.MsgDeleteUserResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.User.Get(ctx, msg.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if err := k.User.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove user")
	}

	return &types.MsgDeleteUserResponse{}, nil
}
