package decon

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "decon/api/decon/decon"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod:      "Sla",
					Use:            "sla [years-as-customer] [number-of-units]",
					Short:          "Query sla",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "yearsAsCustomer"}, {ProtoField: "numberOfUnits"}},
				},

				{
					RpcMethod:      "Fine",
					Use:            "fine [years-as-customer] [number-of-units] [defective-units]",
					Short:          "Query fine",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "yearsAsCustomer"}, {ProtoField: "numberOfUnits"}, {ProtoField: "defectiveUnits"}},
				},

				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
