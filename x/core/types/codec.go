package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
)

func RegisterInterfaces(registrar codectypes.InterfaceRegistry) {
	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRewardMiner{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransfer{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMint{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBlockRecord{},
		&MsgUpdateBlockRecord{},
		&MsgDeleteBlockRecord{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMiner{},
		&MsgUpdateMiner{},
		&MsgDeleteMiner{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateUser{},
		&MsgUpdateUser{},
		&MsgDeleteUser{},
	)

	registrar.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)

	registrar.RegisterImplementations((*txtypes.MsgResponse)(nil),
		&MsgRewardMinerResponse{},
		&MsgTransferResponse{},
		&MsgMintResponse{},
		&MsgCreateBlockRecordResponse{},
		&MsgUpdateBlockRecordResponse{},
		&MsgDeleteBlockRecordResponse{},
		&MsgCreateMinerResponse{},
		&MsgUpdateMinerResponse{},
		&MsgDeleteMinerResponse{},
		&MsgCreateUserResponse{},
		&MsgUpdateUserResponse{},
		&MsgDeleteUserResponse{},
		&MsgUpdateParamsResponse{},
	)
}
