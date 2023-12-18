package finamclient

import "github.com/evsamsonov/FinamTradeGo/v2/tradeapi"

func (f *FinamClient) GetSecurities() (*tradeapi.GetSecuritiesResult, error) {

	res, err := f.securities.GetSecurities(f.ctx, &tradeapi.GetSecuritiesRequest{})
	if err != nil {
		return nil, err
	}
	return res, nil
}
