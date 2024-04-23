package rentchain

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/gopherine/rentchain/testutil/sample"
	rentchainsimulation "github.com/gopherine/rentchain/x/rentchain/simulation"
	"github.com/gopherine/rentchain/x/rentchain/types"
)

// avoid unused import issue
var (
	_ = rentchainsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateRental = "op_weight_msg_create_rental"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRental int = 100

	opWeightMsgCompleteRental = "op_weight_msg_complete_rental"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCompleteRental int = 100

	opWeightMsgCreateRentalAgreement = "op_weight_msg_rental_agreement"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateRentalAgreement int = 100

	opWeightMsgUpdateRentalAgreement = "op_weight_msg_rental_agreement"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateRentalAgreement int = 100

	opWeightMsgDeleteRentalAgreement = "op_weight_msg_rental_agreement"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteRentalAgreement int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	rentchainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		RentalAgreementList: []types.RentalAgreement{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&rentchainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateRental int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateRental, &weightMsgCreateRental, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRental = defaultWeightMsgCreateRental
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRental,
		rentchainsimulation.SimulateMsgCreateRental(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCompleteRental int
	simState.AppParams.GetOrGenerate(opWeightMsgCompleteRental, &weightMsgCompleteRental, nil,
		func(_ *rand.Rand) {
			weightMsgCompleteRental = defaultWeightMsgCompleteRental
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCompleteRental,
		rentchainsimulation.SimulateMsgCompleteRental(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateRentalAgreement int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateRentalAgreement, &weightMsgCreateRentalAgreement, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRentalAgreement = defaultWeightMsgCreateRentalAgreement
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateRentalAgreement,
		rentchainsimulation.SimulateMsgCreateRentalAgreement(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateRentalAgreement int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateRentalAgreement, &weightMsgUpdateRentalAgreement, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateRentalAgreement = defaultWeightMsgUpdateRentalAgreement
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateRentalAgreement,
		rentchainsimulation.SimulateMsgUpdateRentalAgreement(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteRentalAgreement int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteRentalAgreement, &weightMsgDeleteRentalAgreement, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteRentalAgreement = defaultWeightMsgDeleteRentalAgreement
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteRentalAgreement,
		rentchainsimulation.SimulateMsgDeleteRentalAgreement(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateRental,
			defaultWeightMsgCreateRental,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rentchainsimulation.SimulateMsgCreateRental(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCompleteRental,
			defaultWeightMsgCompleteRental,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rentchainsimulation.SimulateMsgCompleteRental(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateRentalAgreement,
			defaultWeightMsgCreateRentalAgreement,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rentchainsimulation.SimulateMsgCreateRentalAgreement(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateRentalAgreement,
			defaultWeightMsgUpdateRentalAgreement,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rentchainsimulation.SimulateMsgUpdateRentalAgreement(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteRentalAgreement,
			defaultWeightMsgDeleteRentalAgreement,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				rentchainsimulation.SimulateMsgDeleteRentalAgreement(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
