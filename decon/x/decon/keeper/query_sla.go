package keeper

import (
	"context"

	"decon/x/decon/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Sla(goCtx context.Context, req *types.QuerySlaRequest) (*types.QuerySlaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	sla := querySla(req.YearsAsCustomer, req.NumberOfUnits)
	return &types.QuerySlaResponse{Sla: sla}, nil
}
