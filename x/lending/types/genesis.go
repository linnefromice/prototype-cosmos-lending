package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PairPoolList: []PairPool{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in pairPool
	pairPoolIdMap := make(map[uint64]bool)
	pairPoolCount := gs.GetPairPoolCount()
	for _, elem := range gs.PairPoolList {
		if _, ok := pairPoolIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for pairPool")
		}
		if elem.Id >= pairPoolCount {
			return fmt.Errorf("pairPool id should be lower or equal than the last id")
		}
		pairPoolIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
