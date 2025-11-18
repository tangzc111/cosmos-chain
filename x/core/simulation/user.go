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

func SimulateMsgCreateUser(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)

		i := r.Int()
		msg := &types.MsgCreateUser{
			Creator: simAccount.Address.String(),
			Index:   strconv.Itoa(i),
		}

		found, err := k.User.Has(ctx, msg.Index)
		if err == nil && found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "User already exist"), nil, nil
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

func SimulateMsgUpdateUser(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			user       = types.User{}
			msg        = &types.MsgUpdateUser{}
			found      = false
		)

		var allUser []types.User
		err := k.User.Walk(ctx, nil, func(key string, value types.User) (stop bool, err error) {
			allUser = append(allUser, value)
			return false, nil
		})
		if err != nil {
			panic(err)
		}

		for _, obj := range allUser {
			acc, err := ak.AddressCodec().StringToBytes(obj.Creator)
			if err != nil {
				return simtypes.OperationMsg{}, nil, err
			}

			simAccount, found = simtypes.FindAccount(accs, sdk.AccAddress(acc))
			if found {
				user = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "user creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.Index = user.Index

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

func SimulateMsgDeleteUser(
	ak types.AuthKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
	txGen client.TxConfig,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		var (
			simAccount = simtypes.Account{}
			user       = types.User{}
			msg        = &types.MsgUpdateUser{}
			found      = false
		)

		var allUser []types.User
		err := k.User.Walk(ctx, nil, func(key string, value types.User) (stop bool, err error) {
			allUser = append(allUser, value)
			return false, nil
		})
		if err != nil {
			panic(err)
		}

		for _, obj := range allUser {
			acc, err := ak.AddressCodec().StringToBytes(obj.Creator)
			if err != nil {
				return simtypes.OperationMsg{}, nil, err
			}

			simAccount, found = simtypes.FindAccount(accs, sdk.AccAddress(acc))
			if found {
				user = obj
				break
			}
		}
		if !found {
			return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "user creator not found"), nil, nil
		}
		msg.Creator = simAccount.Address.String()
		msg.Index = user.Index

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
