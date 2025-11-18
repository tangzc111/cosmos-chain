package types

import "fmt"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:  DefaultParams(),
		UserMap: []User{}, MinerMap: []Miner{}}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	userIndexMap := make(map[string]struct{})

	for _, elem := range gs.UserMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := userIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for user")
		}
		userIndexMap[index] = struct{}{}
	}
	minerIndexMap := make(map[string]struct{})

	for _, elem := range gs.MinerMap {
		index := fmt.Sprint(elem.Index)
		if _, ok := minerIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for miner")
		}
		minerIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
