package finamclient

import "github.com/DBoyara/FinamTradeGo/tradeapi"

func (f *FinamClient) GetSecurities() (*tradeapi.GetSecuritiesResult, error) {
	f.CreateRequestContext()

	res, err := f.securities.GetSecurities(f.ctx, &tradeapi.GetSecuritiesRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
