package keeper

import (
	"context"
	"fmt"

	"tokenchain/x/core/types"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) mustValidAddress(addr string) ([]byte, error) {
	if addr == "" {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidAddress, "address cannot be empty")
	}

	decoded, err := k.addressCodec.StringToBytes(addr)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid bech32 address: %s", err)
	}

	return decoded, nil
}

func parseAmount(amount string) (sdkmath.Int, error) {
	value, ok := sdkmath.NewIntFromString(amount)
	if !ok {
		return sdkmath.Int{}, errorsmod.Wrap(types.ErrInvalidCoin, "invalid amount")
	}

	if !value.IsPositive() {
		return sdkmath.Int{}, errorsmod.Wrap(types.ErrInvalidCoin, "amount must be positive")
	}

	return value, nil
}

func buildCoin(amount sdkmath.Int, denom string) (sdk.Coin, error) {
	if denom == "" {
		return sdk.Coin{}, errorsmod.Wrap(types.ErrInvalidCoin, "denom cannot be empty")
	}

	return sdk.NewCoin(denom, amount), nil
}

func (k Keeper) normalizeUserKey(address string) string {
	return address
}

func (k Keeper) ensureUserExists(ctx context.Context, address string) (types.User, error) {
	key := k.normalizeUserKey(address)
	user, err := k.User.Get(ctx, key)
	if err != nil {
		return types.User{}, errorsmod.Wrap(types.ErrUserNotFound, fmt.Sprintf("address %s", address))
	}

	return user, nil
}

func (k Keeper) ensureMinerExists(ctx context.Context, address string) (types.Miner, error) {
	miner, err := k.Miner.Get(ctx, k.normalizeUserKey(address))
	if err != nil {
		return types.Miner{}, errorsmod.Wrap(types.ErrMinerNotFound, fmt.Sprintf("address %s", address))
	}

	return miner, nil
}
