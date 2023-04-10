package finamclient

import (
	"github.com/DBoyara/FinamTradeGo/tradeapi"
)

func (f *FinamClient) NewOrder(in *tradeapi.NewOrderRequest) (*tradeapi.NewOrderResult, error) {
	f.CreateRequestContext()

	res, err := f.orders.NewOrder(f.ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *FinamClient) CancelOrder(transactionId int32) (*tradeapi.CancelOrderResult, error) {
	f.CreateRequestContext()

	in := &tradeapi.CancelOrderRequest{
		ClientId:      f.clientId,
		TransactionId: transactionId,
	}

	res, err := f.orders.CancelOrder(f.ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (f *FinamClient) GetOrders(includeMatched, includeCanceled, includeActive bool) (*tradeapi.GetOrdersResult, error) {
	f.CreateRequestContext()

	in := &tradeapi.GetOrdersRequest{
		ClientId:        f.clientId,
		IncludeMatched:  includeMatched,
		IncludeCanceled: includeCanceled,
		IncludeActive:   includeActive,
	}

	res, err := f.orders.GetOrders(f.ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
