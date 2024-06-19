package keeper

import (
	"context"
	"github.com/comdex-official/comdex/x/bandoracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) FetchPriceResult(c context.Context, req *types.QueryFetchPriceRequest) (*types.QueryFetchPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	result, err := k.GetFetchPriceResult(ctx, types.OracleRequestID(req.RequestId))
	if err != nil {
		return nil, err
	}
	return &types.QueryFetchPriceResponse{Result: &result}, nil
}

func (k Keeper) LastFetchPriceID(c context.Context, _ *types.QueryLastFetchPriceIdRequest) (*types.QueryLastFetchPriceIdResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	id := k.GetLastFetchPriceID(ctx)
	return &types.QueryLastFetchPriceIdResponse{RequestId: id}, nil
}

func (k Keeper) FetchPriceData(c context.Context, _ *types.QueryFetchPriceDataRequest) (*types.QueryFetchPriceDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	data := k.GetFetchPriceMsg(ctx)
	return &types.QueryFetchPriceDataResponse{MsgFetchPriceData: data}, nil
}

func (k Keeper) DiscardData(c context.Context, _ *types.QueryDiscardDataRequest) (*types.QueryDiscardDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	dd := k.GetDiscardData(ctx)
	return &types.QueryDiscardDataResponse{DiscardData: dd}, nil
}
