package types_test

import (
	"testing"

	"tokenchain/x/core/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				UserMap: []types.User{{Index: "0"}, {Index: "1"}}, MinerMap: []types.Miner{{Index: "0"}, {Index: "1"}}, BlockRecordMap: []types.BlockRecord{{Index: "0"}, {Index: "1"}}},
			valid: true,
		}, {
			desc: "duplicated user",
			genState: &types.GenesisState{
				UserMap: []types.User{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
				MinerMap: []types.Miner{{Index: "0"}, {Index: "1"}}, BlockRecordMap: []types.BlockRecord{{Index: "0"}, {Index: "1"}}},
			valid: false,
		}, {
			desc: "duplicated miner",
			genState: &types.GenesisState{
				MinerMap: []types.Miner{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
				BlockRecordMap: []types.BlockRecord{{Index: "0"}, {Index: "1"}}},
			valid: false,
		}, {
			desc: "duplicated blockRecord",
			genState: &types.GenesisState{
				BlockRecordMap: []types.BlockRecord{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
