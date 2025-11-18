package keeper

import (
	"encoding/hex"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"cosmos-chain/x/core/types"
)

// BeginBlocker records basic metadata about every committed block so it can be
// queried later through the module API.
func (k Keeper) BeginBlocker(ctx sdk.Context) error {
	header := ctx.BlockHeader()
	heightStr := strconv.FormatInt(header.Height, 10)

	record := types.BlockRecord{
		Creator:  types.ModuleName,
		Index:    heightStr,
		Height:   heightStr,
		Hash:     hex.EncodeToString(header.AppHash),
		Proposer: hex.EncodeToString(header.ProposerAddress),
		Time:     header.Time.UTC().Format(time.RFC3339),
	}

	if record.Hash == "" && header.LastBlockId.Hash != nil {
		record.Hash = hex.EncodeToString(header.LastBlockId.Hash)
	}

	if err := k.BlockRecord.Set(ctx, record.Index, record); err != nil {
		return err
	}

	return nil
}
