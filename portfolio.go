package finamclient

import "github.com/DBoyara/FinamTradeGo/tradeapi"

func (f *FinamClient) GetPortfolio(includeCurrencies, includeMoney, includePositions, includeMaxBuySell bool) (*tradeapi.GetPortfolioResult, error) {

	in := &tradeapi.GetPortfolioRequest{
		ClientId: f.clientId,
		Content: &tradeapi.PortfolioContent{
			IncludeCurrencies: includeCurrencies,
			IncludeMoney:      includeMoney,
			IncludePositions:  includePositions,
			IncludeMaxBuySell: includeMaxBuySell,
		},
	}

	res, err := f.portfolio.GetPortfolio(f.ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
