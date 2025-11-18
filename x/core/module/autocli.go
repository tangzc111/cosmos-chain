package core

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	"tokenchain/x/core/types"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: types.Query_serviceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ListUser",
					Use:       "list-user",
					Short:     "List all user",
				},
				{
					RpcMethod:      "GetUser",
					Use:            "get-user [id]",
					Short:          "Gets a user",
					Alias:          []string{"show-user"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              types.Msg_serviceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateUser",
					Use:            "create-user [index] [address] [username] [description]",
					Short:          "Create a new user",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "address"}, {ProtoField: "username"}, {ProtoField: "description"}},
				},
				{
					RpcMethod:      "UpdateUser",
					Use:            "update-user [index] [address] [username] [description]",
					Short:          "Update user",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "address"}, {ProtoField: "username"}, {ProtoField: "description"}},
				},
				{
					RpcMethod:      "DeleteUser",
					Use:            "delete-user [index]",
					Short:          "Delete user",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
