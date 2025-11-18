package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/core module sentinel errors
var (
	ErrInvalidSigner   = errors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrUserExists      = errors.Register(ModuleName, 1101, "user already exists")
	ErrUserNotFound    = errors.Register(ModuleName, 1102, "user not found")
	ErrMinerExists     = errors.Register(ModuleName, 1103, "miner already exists")
	ErrMinerNotFound   = errors.Register(ModuleName, 1104, "miner not found")
	ErrInvalidCoin     = errors.Register(ModuleName, 1105, "invalid coin value")
	ErrUnauthorized    = errors.Register(ModuleName, 1106, "unauthorized")
	ErrBlockNotTracked = errors.Register(ModuleName, 1107, "block not tracked yet")
)
