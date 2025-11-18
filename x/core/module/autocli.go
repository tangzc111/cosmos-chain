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
				{
					RpcMethod: "ListMiner",
					Use:       "list-miner",
					Short:     "List all miner",
				},
				{
					RpcMethod:      "GetMiner",
					Use:            "get-miner [id]",
					Short:          "Gets a miner",
					Alias:          []string{"show-miner"},
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod: "ListBlockRecord",
					Use:       "list-block-record",
					Short:     "List all block-record",
				},
				{
					RpcMethod:      "GetBlockRecord",
					Use:            "get-block-record [id]",
					Short:          "Gets a block-record",
					Alias:          []string{"show-block-record"},
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
				{
					RpcMethod:      "CreateMiner",
					Use:            "create-miner [index] [address] [power] [description] [total-reward]",
					Short:          "Create a new miner",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "address"}, {ProtoField: "power"}, {ProtoField: "description"}, {ProtoField: "total_reward"}},
				},
				{
					RpcMethod:      "UpdateMiner",
					Use:            "update-miner [index] [address] [power] [description] [total-reward]",
					Short:          "Update miner",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "address"}, {ProtoField: "power"}, {ProtoField: "description"}, {ProtoField: "total_reward"}},
				},
				{
					RpcMethod:      "DeleteMiner",
					Use:            "delete-miner [index]",
					Short:          "Delete miner",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "CreateBlockRecord",
					Use:            "create-block-record [index] [height] [hash] [proposer] [time]",
					Short:          "Create a new block-record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "height"}, {ProtoField: "hash"}, {ProtoField: "proposer"}, {ProtoField: "time"}},
				},
				{
					RpcMethod:      "UpdateBlockRecord",
					Use:            "update-block-record [index] [height] [hash] [proposer] [time]",
					Short:          "Update block-record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}, {ProtoField: "height"}, {ProtoField: "hash"}, {ProtoField: "proposer"}, {ProtoField: "time"}},
				},
				{
					RpcMethod:      "DeleteBlockRecord",
					Use:            "delete-block-record [index]",
					Short:          "Delete block-record",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "index"}},
				},
				{
					RpcMethod:      "Mint",
					Use:            "mint [recipient] [amount] [denom]",
					Short:          "Send a mint tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "recipient"}, {ProtoField: "amount"}, {ProtoField: "denom"}},
				},
				{
			RpcMethod: "Transfer",
			Use: "transfer [to] [amount] [denom]",
			Short: "Send a transfer tx",
			PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "to"}, {ProtoField: "amount"}, {ProtoField: "denom"}},
		},
		// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
