package keeper

import (
	"context"

	"tokenchain/x/core/types"

	errorsmod "cosmossdk.io/errors"
)

func (k msgServer) RewardMiner(ctx context.Context, msg *types.MsgRewardMiner) (*types.MsgRewardMinerResponse, error) {
	if _, err := k.addressCodec.StringToBytes(msg.Creator); err != nil {
		return nil, errorsmod.Wrap(err, "invalid authority address")
	}

	// TODO: Handle the message

	return &types.MsgRewardMinerResponse{}, nil
}
