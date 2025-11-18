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

func (k msgServer) CreateMiner(ctx context.Context, msg *types.MsgCreateMiner) (*types.MsgCreateMinerResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.Miner.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var miner = types.Miner{
		Creator:     msg.Creator,
		Index:       msg.Index,
		Address:     msg.Address,
		Power:       msg.Power,
		Description: msg.Description,
		TotalReward: msg.TotalReward,
	}

	if err := k.Miner.Set(ctx, miner.Index, miner); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateMinerResponse{}, nil
}

func (k msgServer) UpdateMiner(ctx context.Context, msg *types.MsgUpdateMiner) (*types.MsgUpdateMinerResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Miner.Get(ctx, msg.Index)
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

	var miner = types.Miner{
		Creator:     msg.Creator,
		Index:       msg.Index,
		Address:     msg.Address,
		Power:       msg.Power,
		Description: msg.Description,
		TotalReward: msg.TotalReward,
	}

	if err := k.Miner.Set(ctx, miner.Index, miner); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update miner")
	}

	return &types.MsgUpdateMinerResponse{}, nil
}

func (k msgServer) DeleteMiner(ctx context.Context, msg *types.MsgDeleteMiner) (*types.MsgDeleteMinerResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.Miner.Get(ctx, msg.Index)
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

	if err := k.Miner.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove miner")
	}

	return &types.MsgDeleteMinerResponse{}, nil
}
