package keeper

import (
	"context"

	"decon/x/decon/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Fine(goCtx context.Context, req *types.QueryFineRequest) (*types.QueryFineResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	fine := queryFine(req.YearsAsCustomer, req.NumberOfUnits, req.DefectiveUnits)
	return &types.QueryFineResponse{Fine: fine}, nil
}
