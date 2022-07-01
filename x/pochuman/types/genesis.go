package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		FeeBalanceList:      []FeeBalance{},
		KeysignVoteDataList: []KeysignVoteData{},
		ObserveVoteList:     []ObserveVote{},
		PoolBalanceList:     []PoolBalance{},
		TransactionDataList: []TransactionData{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in feeBalance
	feeBalanceIndexMap := make(map[string]struct{})

	for _, elem := range gs.FeeBalanceList {
		index := string(FeeBalanceKey(elem.Index))
		if _, ok := feeBalanceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for feeBalance")
		}
		feeBalanceIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in keysignVoteData
	keysignVoteDataIndexMap := make(map[string]struct{})

	for _, elem := range gs.KeysignVoteDataList {
		index := string(KeysignVoteDataKey(elem.Index))
		if _, ok := keysignVoteDataIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for keysignVoteData")
		}
		keysignVoteDataIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in observeVote
	observeVoteIndexMap := make(map[string]struct{})

	for _, elem := range gs.ObserveVoteList {
		index := string(ObserveVoteKey(elem.Index))
		if _, ok := observeVoteIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for observeVote")
		}
		observeVoteIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in poolBalance
	poolBalanceIndexMap := make(map[string]struct{})

	for _, elem := range gs.PoolBalanceList {
		index := string(PoolBalanceKey(elem.Index))
		if _, ok := poolBalanceIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for poolBalance")
		}
		poolBalanceIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in transactionData
	transactionDataIndexMap := make(map[string]struct{})

	for _, elem := range gs.TransactionDataList {
		index := string(TransactionDataKey(elem.Index))
		if _, ok := transactionDataIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for transactionData")
		}
		transactionDataIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
