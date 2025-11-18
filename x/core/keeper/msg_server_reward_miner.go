package keeper

import (
	"context"

	"tokenchain/x/core/types"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RewardMiner(ctx context.Context, msg *types.MsgRewardMiner) (*types.MsgRewardMinerResponse, error) {
	if _, err := k.mustValidAddress(msg.Creator); err != nil {
		return nil, err
	}
	if _, err := k.ensureUserExists(ctx, msg.Creator); err != nil {
		return nil, err
	}

	minerAddrBz, err := k.mustValidAddress(msg.Miner)
	if err != nil {
		return nil, err
	}

	miner, err := k.ensureMinerExists(ctx, msg.Miner)
	if err != nil {
		return nil, err
	}

	amount, err := parseAmount(msg.Amount)
	if err != nil {
		return nil, err
	}

	coin, err := buildCoin(amount, msg.Denom)
	if err != nil {
		return nil, err
	}

	coins := sdk.NewCoins(coin)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, coins); err != nil {
		return nil, errorsmod.Wrap(err, "mint coins failed")
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(minerAddrBz), coins); err != nil {
		return nil, errorsmod.Wrap(err, "reward transfer failed")
	}

	current, ok := sdkmath.NewIntFromString(miner.TotalReward)
	if !ok {
		current = sdkmath.ZeroInt()
	}

	miner.TotalReward = current.Add(amount).String()
	if err := k.Miner.Set(ctx, miner.Index, miner); err != nil {
		return nil, errorsmod.Wrap(err, "update miner reward failed")
	}

	return &types.MsgRewardMinerResponse{}, nil
}
