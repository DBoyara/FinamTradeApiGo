package finamclient

import "github.com/DBoyara/FinamTradeGo/tradeapi"

func (f *FinamClient) GetDayCandles(in *tradeapi.GetDayCandlesRequest) (*tradeapi.GetDayCandlesResult, error) {
	res, err := f.candles.GetDayCandles(f.ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *FinamClient) GetIntradayCandles(in *tradeapi.GetIntradayCandlesRequest) (*tradeapi.GetIntradayCandlesResult, error) {

	res, err := f.candles.GetIntradayCandles(f.ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
