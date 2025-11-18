package keeper

import (
	"context"

	"tokenchain/x/core/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Transfer(ctx context.Context, msg *types.MsgTransfer) (*types.MsgTransferResponse, error) {
	fromBz, err := k.mustValidAddress(msg.Creator)
	if err != nil {
		return nil, err
	}

	toBz, err := k.mustValidAddress(msg.To)
	if err != nil {
		return nil, err
	}

	if _, err := k.ensureUserExists(ctx, msg.Creator); err != nil {
		return nil, err
	}
	if _, err := k.ensureUserExists(ctx, msg.To); err != nil {
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

	if err := k.bankKeeper.SendCoins(ctx, sdk.AccAddress(fromBz), sdk.AccAddress(toBz), sdk.NewCoins(coin)); err != nil {
		return nil, errorsmod.Wrap(err, "transfer failed")
	}

	return &types.MsgTransferResponse{}, nil
}
