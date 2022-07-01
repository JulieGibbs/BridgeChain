package types_test

import (
	"testing"

	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
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

				FeeBalanceList: []types.FeeBalance{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				KeysignVoteDataList: []types.KeysignVoteData{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				ObserveVoteList: []types.ObserveVote{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				PoolBalanceList: []types.PoolBalance{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				TransactionDataList: []types.TransactionData{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated feeBalance",
			genState: &types.GenesisState{
				FeeBalanceList: []types.FeeBalance{
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
		{
			desc: "duplicated keysignVoteData",
			genState: &types.GenesisState{
				KeysignVoteDataList: []types.KeysignVoteData{
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
		{
			desc: "duplicated observeVote",
			genState: &types.GenesisState{
				ObserveVoteList: []types.ObserveVote{
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
		{
			desc: "duplicated poolBalance",
			genState: &types.GenesisState{
				PoolBalanceList: []types.PoolBalance{
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
		{
			desc: "duplicated transactionData",
			genState: &types.GenesisState{
				TransactionDataList: []types.TransactionData{
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
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
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
