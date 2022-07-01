package pochuman_test

import (
	"testing"

	keepertest "github.com/VigorousDeveloper/poc-human/testutil/keeper"
	"github.com/VigorousDeveloper/poc-human/testutil/nullify"
	"github.com/VigorousDeveloper/poc-human/x/pochuman"
	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PochumanKeeper(t)
	pochuman.InitGenesis(ctx, *k, genesisState)
	got := pochuman.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.FeeBalanceList, got.FeeBalanceList)
	require.ElementsMatch(t, genesisState.KeysignVoteDataList, got.KeysignVoteDataList)
	require.ElementsMatch(t, genesisState.ObserveVoteList, got.ObserveVoteList)
	require.ElementsMatch(t, genesisState.PoolBalanceList, got.PoolBalanceList)
	require.ElementsMatch(t, genesisState.TransactionDataList, got.TransactionDataList)
	// this line is used by starport scaffolding # genesis/test/assert
}
