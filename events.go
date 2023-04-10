package finamclient

import (
	"time"

	"github.com/DBoyara/FinamTradeGo/tradeapi"
)

func (f *FinamClient) GetOrderBooksChan() chan *tradeapi.OrderBookEvent {
	return f.orderBooksChan
}

func (f *FinamClient) GetOrderTradeChan() chan *tradeapi.TradeEvent {
	return f.orderTradeChan
}

func (f *FinamClient) GetOrderChan() chan *tradeapi.OrderEvent {
	return f.orderChan
}

func (f *FinamClient) GetErrorChan() chan error {
	return f.errChan
}

func (f *FinamClient) SubscribeOrderBook(in *tradeapi.OrderBookSubscribeRequest) {
	f.CreateRequestContext()

	e, err := f.events.GetEvents(f.ctx)
	if err != nil {
		f.errChan <- err
	}

	payload := &tradeapi.SubscriptionRequest{
		Payload: &tradeapi.SubscriptionRequest_OrderBookSubscribeRequest{
			OrderBookSubscribeRequest: in,
		},
	}

	err = e.Send(payload)
	if err != nil {
		f.errChan <- err
	}

	for {
		msg, err := e.Recv()
		if err != nil {
			f.errChan <- err
		}

		f.orderBooksChan <- msg.GetOrderBook()
		time.Sleep(1 * time.Second)
	}
}

func (f *FinamClient) UnSubscribeOrderBook(in *tradeapi.OrderBookUnsubscribeRequest) *tradeapi.ResponseEvent {
	f.CreateRequestContext()

	e, err := f.events.GetEvents(f.ctx)
	if err != nil {
		f.errChan <- err
	}

	payload := &tradeapi.SubscriptionRequest{
		Payload: &tradeapi.SubscriptionRequest_OrderBookUnsubscribeRequest{
			OrderBookUnsubscribeRequest: in,
		},
	}

	err = e.Send(payload)
	if err != nil {
		f.errChan <- err
	}

	msg, err := e.Recv()
	if err != nil {
		f.errChan <- err
	}

	return msg.GetResponse()
}

func (f *FinamClient) SubscribeOrderTrade(in *tradeapi.OrderTradeSubscribeRequest) {
	f.CreateRequestContext()

	e, err := f.events.GetEvents(f.ctx)
	if err != nil {
		f.errChan <- err
	}

	payload := &tradeapi.SubscriptionRequest{
		Payload: &tradeapi.SubscriptionRequest_OrderTradeSubscribeRequest{
			OrderTradeSubscribeRequest: in,
		},
	}

	err = e.Send(payload)
	if err != nil {
		f.errChan <- err
	}

	for {
		msg, err := e.Recv()
		if err != nil {
			f.errChan <- err
		}

		f.orderTradeChan <- msg.GetTrade()
		f.orderChan <- msg.GetOrder()
		time.Sleep(1 * time.Second)
	}
}

func (f *FinamClient) UnSubscribeOrderTrade(in *tradeapi.OrderTradeUnsubscribeRequest) *tradeapi.ResponseEvent {
	f.CreateRequestContext()

	e, err := f.events.GetEvents(f.ctx)
	if err != nil {
		f.errChan <- err
	}

	payload := &tradeapi.SubscriptionRequest{
		Payload: &tradeapi.SubscriptionRequest_OrderTradeUnsubscribeRequest{
			OrderTradeUnsubscribeRequest: in,
		},
	}

	err = e.Send(payload)
	if err != nil {
		f.errChan <- err
	}

	msg, err := e.Recv()
	if err != nil {
		f.errChan <- err
	}

	return msg.GetResponse()
}
