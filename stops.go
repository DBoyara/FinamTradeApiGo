package finamclient

import (
	"github.com/evsamsonov/FinamTradeGo/v2/tradeapi"
)

func (f *FinamClient) NewStop(in *tradeapi.NewStopRequest) (*tradeapi.NewStopResult, error) {

	res, err := f.stops.NewStop(f.ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *FinamClient) CancelStop(stopId int32) (*tradeapi.CancelStopResult, error) {

	in := &tradeapi.CancelStopRequest{
		ClientId: f.clientId,
		StopId:   stopId,
	}

	res, err := f.stops.CancelStop(f.ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *FinamClient) GetStops(IncludeExecuted, includeCanceled, includeActive bool) (*tradeapi.GetStopsResult, error) {

	in := &tradeapi.GetStopsRequest{
		ClientId:        f.clientId,
		IncludeExecuted: IncludeExecuted,
		IncludeCanceled: includeCanceled,
		IncludeActive:   includeActive,
	}

	res, err := f.stops.GetStops(f.ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
