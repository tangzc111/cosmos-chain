package keeper

import (
	"context"
	"errors"
	"strconv"

	"tokenchain/x/core/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateMiner(ctx context.Context, msg *types.MsgCreateMiner) (*types.MsgCreateMinerResponse, error) {
	if _, err := k.mustValidAddress(msg.Creator); err != nil {
		return nil, err
	}
	if _, err := k.mustValidAddress(msg.Address); err != nil {
		return nil, err
	}
	if msg.Index != "" && msg.Index != msg.Address {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index must match miner address")
	}
	if msg.Creator != msg.Address {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "creator must match miner address")
	}
	if _, err := k.ensureUserExists(ctx, msg.Address); err != nil {
		return nil, err
	}

	if _, err := strconv.ParseUint(msg.Power, 10, 64); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "power must be a positive integer")
	}

	key := k.normalizeUserKey(msg.Address)

	// Check if the value already exists
	ok, err := k.Miner.Has(ctx, key)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, types.ErrMinerExists
	}

	var miner = types.Miner{
		Creator:     msg.Creator,
		Index:       key,
		Address:     msg.Address,
		Power:       msg.Power,
		Description: msg.Description,
		TotalReward: "0",
	}

	if err := k.Miner.Set(ctx, key, miner); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateMinerResponse{}, nil
}

func (k msgServer) UpdateMiner(ctx context.Context, msg *types.MsgUpdateMiner) (*types.MsgUpdateMinerResponse, error) {
	if _, err := k.mustValidAddress(msg.Creator); err != nil {
		return nil, err
	}

	key := k.normalizeUserKey(msg.Index)

	// Check if the value exists
	val, err := k.Miner.Get(ctx, key)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, types.ErrMinerNotFound
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

	if msg.Power != "" {
		if _, err := strconv.ParseUint(msg.Power, 10, 64); err != nil {
			return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "power must be a positive integer")
		}
		val.Power = msg.Power
	}

	if msg.Description != "" {
		val.Description = msg.Description
	}

	if err := k.Miner.Set(ctx, val.Index, val); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update miner")
	}

	return &types.MsgUpdateMinerResponse{}, nil
}

func (k msgServer) DeleteMiner(ctx context.Context, msg *types.MsgDeleteMiner) (*types.MsgDeleteMinerResponse, error) {
	if _, err := k.mustValidAddress(msg.Creator); err != nil {
		return nil, err
	}

	// Check if the value exists
	val, err := k.Miner.Get(ctx, k.normalizeUserKey(msg.Index))
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, types.ErrMinerNotFound
		}

		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(types.ErrUnauthorized, "incorrect owner")
	}

	if err := k.Miner.Remove(ctx, val.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove miner")
	}

	return &types.MsgDeleteMinerResponse{}, nil
}
