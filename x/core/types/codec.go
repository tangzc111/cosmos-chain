package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterInterfaces(registrar codectypes.InterfaceRegistry) {
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
	msgservice.RegisterMsgServiceDesc(registrar, &_Msg_serviceDesc)
}
