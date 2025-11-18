package keeper

import (
	"context"

	"tokenchain/x/core/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Mint(ctx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	if _, err := k.mustValidAddress(msg.Creator); err != nil {
		return nil, err
	}
	if _, err := k.ensureUserExists(ctx, msg.Creator); err != nil {
		return nil, err
	}

	recipientBz, err := k.mustValidAddress(msg.Recipient)
	if err != nil {
		return nil, err
	}
	if _, err := k.ensureUserExists(ctx, msg.Recipient); err != nil {
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

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(recipientBz), coins); err != nil {
		return nil, errorsmod.Wrap(err, "transfer minted coins failed")
	}

	return &types.MsgMintResponse{}, nil
}
