package core

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"math/rand"

	"tokenchain/testutil/sample"
	coresimulation "tokenchain/x/core/simulation"
	"tokenchain/x/core/types"
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	coreGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		UserMap: []types.User{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}, MinerMap: []types.Miner{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}, BlockRecordMap: []types.BlockRecord{{Creator: sample.AccAddress(),
			Index: "0",
		}, {Creator: sample.AccAddress(),
			Index: "1",
		}}}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&coreGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)
	const (
		opWeightMsgCreateUser          = "op_weight_msg_core"
		defaultWeightMsgCreateUser int = 100
	)

	var weightMsgCreateUser int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateUser, &weightMsgCreateUser, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUser = defaultWeightMsgCreateUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateUser,
		coresimulation.SimulateMsgCreateUser(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateUser          = "op_weight_msg_core"
		defaultWeightMsgUpdateUser int = 100
	)

	var weightMsgUpdateUser int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateUser, &weightMsgUpdateUser, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateUser = defaultWeightMsgUpdateUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateUser,
		coresimulation.SimulateMsgUpdateUser(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteUser          = "op_weight_msg_core"
		defaultWeightMsgDeleteUser int = 100
	)

	var weightMsgDeleteUser int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteUser, &weightMsgDeleteUser, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteUser = defaultWeightMsgDeleteUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteUser,
		coresimulation.SimulateMsgDeleteUser(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreateMiner          = "op_weight_msg_core"
		defaultWeightMsgCreateMiner int = 100
	)

	var weightMsgCreateMiner int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateMiner, &weightMsgCreateMiner, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMiner = defaultWeightMsgCreateMiner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMiner,
		coresimulation.SimulateMsgCreateMiner(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateMiner          = "op_weight_msg_core"
		defaultWeightMsgUpdateMiner int = 100
	)

	var weightMsgUpdateMiner int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateMiner, &weightMsgUpdateMiner, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMiner = defaultWeightMsgUpdateMiner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMiner,
		coresimulation.SimulateMsgUpdateMiner(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteMiner          = "op_weight_msg_core"
		defaultWeightMsgDeleteMiner int = 100
	)

	var weightMsgDeleteMiner int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteMiner, &weightMsgDeleteMiner, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMiner = defaultWeightMsgDeleteMiner
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMiner,
		coresimulation.SimulateMsgDeleteMiner(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgCreateBlockRecord          = "op_weight_msg_core"
		defaultWeightMsgCreateBlockRecord int = 100
	)

	var weightMsgCreateBlockRecord int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateBlockRecord, &weightMsgCreateBlockRecord, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBlockRecord = defaultWeightMsgCreateBlockRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateBlockRecord,
		coresimulation.SimulateMsgCreateBlockRecord(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgUpdateBlockRecord          = "op_weight_msg_core"
		defaultWeightMsgUpdateBlockRecord int = 100
	)

	var weightMsgUpdateBlockRecord int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateBlockRecord, &weightMsgUpdateBlockRecord, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBlockRecord = defaultWeightMsgUpdateBlockRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateBlockRecord,
		coresimulation.SimulateMsgUpdateBlockRecord(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgDeleteBlockRecord          = "op_weight_msg_core"
		defaultWeightMsgDeleteBlockRecord int = 100
	)

	var weightMsgDeleteBlockRecord int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteBlockRecord, &weightMsgDeleteBlockRecord, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteBlockRecord = defaultWeightMsgDeleteBlockRecord
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteBlockRecord,
		coresimulation.SimulateMsgDeleteBlockRecord(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgMint          = "op_weight_msg_core"
		defaultWeightMsgMint int = 100
	)

	var weightMsgMint int
	simState.AppParams.GetOrGenerate(opWeightMsgMint, &weightMsgMint, nil,
		func(_ *rand.Rand) {
			weightMsgMint = defaultWeightMsgMint
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMint,
		coresimulation.SimulateMsgMint(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))
	const (
		opWeightMsgTransfer          = "op_weight_msg_core"
		defaultWeightMsgTransfer int = 100
	)

	var weightMsgTransfer int
	simState.AppParams.GetOrGenerate(opWeightMsgTransfer, &weightMsgTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgTransfer = defaultWeightMsgTransfer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTransfer,
		coresimulation.SimulateMsgTransfer(am.authKeeper, am.bankKeeper, am.keeper, simState.TxConfig),
	))

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{}
}