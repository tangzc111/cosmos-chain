package simulation

import (
	"math/rand"
	"strconv"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"cosmos-chain/x/core/keeper"
	"cosmos-chain/x/core/types"
)

func SimulateMsgCreateBlockRecord(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateBlockRecord{
			Creator: simAccount.Address.String(),
			Index:   strconv.Itoa(i),
		}

		found, err := k.BlockRecord.Has(ctx, msg.Index)
		if err == nil && found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "BlockRecord already exist"), nil, nil
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgUpdateBlockRecord(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount  = simtypes.Account{}
			blockRecord = types.BlockRecord{}
			msg         = &types.MsgUpdateBlockRecord{}
			found       = false
		)

		var allBlockRecord []types.BlockRecord
		err := k.BlockRecord.Walk(ctx, nil, func(key string, value types.BlockRecord) (stop bool, err error) {
			allBlockRecord = append(allBlockRecord, value)
			return false, nil
		})
		if err != nil {
			panic(err)
		}

		for _, obj := range allBlockRecord {
			acc, err := ak.AddressCodec().StringToBytes(obj.Creator)
			if err != nil {
				return simtypes.OperationMsg{}, nil, err
			}

			simAccount, found = simtypes.FindAccount(accs, sdk.AccAddress(acc))
			if found {
				blockRecord = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "blockRecord creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.Index = blockRecord.Index

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func SimulateMsgDeleteBlockRecord(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount  = simtypes.Account{}
			blockRecord = types.BlockRecord{}
			msg         = &types.MsgUpdateBlockRecord{}
			found       = false
		)

		var allBlockRecord []types.BlockRecord
		err := k.BlockRecord.Walk(ctx, nil, func(key string, value types.BlockRecord) (stop bool, err error) {
			allBlockRecord = append(allBlockRecord, value)
			return false, nil
		})
		if err != nil {
			panic(err)
		}

		for _, obj := range allBlockRecord {
			acc, err := ak.AddressCodec().StringToBytes(obj.Creator)
			if err != nil {
				return simtypes.OperationMsg{}, nil, err
			}

			simAccount, found = simtypes.FindAccount(accs, sdk.AccAddress(acc))
			if found {
				blockRecord = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "blockRecord creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.Index = blockRecord.Index

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           txGen,
			Cdc:             nil,
			Msg:             msg,
			Context:         ctx,
			SimAccount:      simAccount,
			ModuleName:      types.ModuleName,
			CoinsSpentInMsg: sdk.NewCoins(),
			AccountKeeper:   ak,
			Bankkeeper:      bk,
		}
		return simulation.GenAndDeliverTxWithRandFees(txCtx)
	}
}
