package pochuman

import (
	"math/rand"

	"github.com/VigorousDeveloper/poc-human/testutil/sample"
	pochumansimulation "github.com/VigorousDeveloper/poc-human/x/pochuman/simulation"
	"github.com/VigorousDeveloper/poc-human/x/pochuman/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = pochumansimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRequestTransaction = "op_weight_msg_request_transaction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestTransaction int = 100

	opWeightMsgObservationVote = "op_weight_msg_observation_vote"
	// TODO: Determine the simulation weight value
	defaultWeightMsgObservationVote int = 100

	opWeightMsgUpdateBalance = "op_weight_msg_update_balance"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateBalance int = 100

	opWeightMsgKeysignVote = "op_weight_msg_keysign_vote"
	// TODO: Determine the simulation weight value
	defaultWeightMsgKeysignVote int = 100

	opWeightMsgApproveTransaction = "op_weight_msg_approve_transaction"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveTransaction int = 100

	opWeightMsgTranfserPoolcoin = "op_weight_msg_tranfser_poolcoin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTranfserPoolcoin int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	pochumanGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&pochumanGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRequestTransaction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRequestTransaction, &weightMsgRequestTransaction, nil,
		func(_ *rand.Rand) {
			weightMsgRequestTransaction = defaultWeightMsgRequestTransaction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestTransaction,
		pochumansimulation.SimulateMsgRequestTransaction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgObservationVote int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgObservationVote, &weightMsgObservationVote, nil,
		func(_ *rand.Rand) {
			weightMsgObservationVote = defaultWeightMsgObservationVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgObservationVote,
		pochumansimulation.SimulateMsgObservationVote(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateBalance int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateBalance, &weightMsgUpdateBalance, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBalance = defaultWeightMsgUpdateBalance
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateBalance,
		pochumansimulation.SimulateMsgUpdateBalance(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgKeysignVote int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgKeysignVote, &weightMsgKeysignVote, nil,
		func(_ *rand.Rand) {
			weightMsgKeysignVote = defaultWeightMsgKeysignVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgKeysignVote,
		pochumansimulation.SimulateMsgKeysignVote(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveTransaction int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveTransaction, &weightMsgApproveTransaction, nil,
		func(_ *rand.Rand) {
			weightMsgApproveTransaction = defaultWeightMsgApproveTransaction
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveTransaction,
		pochumansimulation.SimulateMsgApproveTransaction(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTranfserPoolcoin int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTranfserPoolcoin, &weightMsgTranfserPoolcoin, nil,
		func(_ *rand.Rand) {
			weightMsgTranfserPoolcoin = defaultWeightMsgTranfserPoolcoin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTranfserPoolcoin,
		pochumansimulation.SimulateMsgTranfserPoolcoin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
