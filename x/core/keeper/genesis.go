package keeper

import (
	"context"

	"cosmos-chain/x/core/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx context.Context, genState types.GenesisState) error {
	for _, elem := range genState.UserMap {
		if err := k.User.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}
	for _, elem := range genState.MinerMap {
		if err := k.Miner.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}
	for _, elem := range genState.BlockRecordMap {
		if err := k.BlockRecord.Set(ctx, elem.Index, elem); err != nil {
			return err
		}
	}

	return k.Params.Set(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func (k Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error) {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}
	if err := k.User.Walk(ctx, nil, func(_ string, val types.User) (stop bool, err error) {
		genesis.UserMap = append(genesis.UserMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	if err := k.Miner.Walk(ctx, nil, func(_ string, val types.Miner) (stop bool, err error) {
		genesis.MinerMap = append(genesis.MinerMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}
	if err := k.BlockRecord.Walk(ctx, nil, func(_ string, val types.BlockRecord) (stop bool, err error) {
		genesis.BlockRecordMap = append(genesis.BlockRecordMap, val)
		return false, nil
	}); err != nil {
		return nil, err
	}

	return genesis, nil
}
