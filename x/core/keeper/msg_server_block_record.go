package keeper

import (
	"context"
	"errors"
	"fmt"

	"cosmos-chain/x/core/types"

	"cosmossdk.io/collections"
	errorsmod "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateBlockRecord(ctx context.Context, msg *types.MsgCreateBlockRecord) (*types.MsgCreateBlockRecordResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid address: %s", err))
	}

	// Check if the value already exists
	ok, err := k.BlockRecord.Has(ctx, msg.Index)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	} else if ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var blockRecord = types.BlockRecord{
		Creator:  msg.Creator,
		Index:    msg.Index,
		Height:   msg.Height,
		Hash:     msg.Hash,
		Proposer: msg.Proposer,
		Time:     msg.Time,
	}

	if err := k.BlockRecord.Set(ctx, blockRecord.Index, blockRecord); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	return &types.MsgCreateBlockRecordResponse{}, nil
}

func (k msgServer) UpdateBlockRecord(ctx context.Context, msg *types.MsgUpdateBlockRecord) (*types.MsgUpdateBlockRecordResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.BlockRecord.Get(ctx, msg.Index)
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

	var blockRecord = types.BlockRecord{
		Creator:  msg.Creator,
		Index:    msg.Index,
		Height:   msg.Height,
		Hash:     msg.Hash,
		Proposer: msg.Proposer,
		Time:     msg.Time,
	}

	if err := k.BlockRecord.Set(ctx, blockRecord.Index, blockRecord); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to update blockRecord")
	}

	return &types.MsgUpdateBlockRecordResponse{}, nil
}

func (k msgServer) DeleteBlockRecord(ctx context.Context, msg *types.MsgDeleteBlockRecord) (*types.MsgDeleteBlockRecordResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, fmt.Sprintf("invalid signer address: %s", err))
	}

	// Check if the value exists
	val, err := k.BlockRecord.Get(ctx, msg.Index)
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

	if err := k.BlockRecord.Remove(ctx, msg.Index); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "failed to remove blockRecord")
	}

	return &types.MsgDeleteBlockRecordResponse{}, nil
}
