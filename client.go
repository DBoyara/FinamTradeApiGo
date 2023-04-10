package finamclient

import (
	"context"
	"crypto/tls"

	"github.com/DBoyara/FinamTradeGo/tradeapi"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	grpcMetadata "google.golang.org/grpc/metadata"
)

type IFinamClient interface {
	// Подписка на биржевой стакан
	SubscribeOrderBook(in *tradeapi.OrderBookSubscribeRequest)
	// Удаление подписки на биржевой стакан
	UnSubscribeOrderBook(in *tradeapi.OrderBookUnsubscribeRequest) *tradeapi.ResponseEvent
	// Подписка на заявки и сделки
	SubscribeOrderTrade(in *tradeapi.OrderTradeSubscribeRequest)
	// Удаление подписки на заявки и сделки
	UnSubscribeOrderTrade(in *tradeapi.OrderTradeUnsubscribeRequest) *tradeapi.ResponseEvent
	// Создать новую заявку.
	NewOrder(in *tradeapi.NewOrderRequest) (*tradeapi.NewOrderResult, error)
	// Отменяет заявку.
	CancelOrder(transactionId int32) (*tradeapi.CancelOrderResult, error)
	// Возвращает список заявок.
	GetOrders(includeMatched, includeCanceled, includeActive bool) (*tradeapi.GetOrdersResult, error)
	// Возвращает список стоп-заявок.
	GetStops(includeExecuted, includeCanceled, includeActive bool) (*tradeapi.GetStopsResult, error)
	// Снимает стоп-заявку.
	CancelStop(stopId int32) (*tradeapi.CancelStopResult, error)
	// Выставляет стоп-заявку.
	NewStop(in *tradeapi.NewStopRequest) (*tradeapi.NewStopResult, error)
	// Посмотреть портфель
	GetPortfolio(includeCurrencies, includeMoney, includePositions, includeMaxBuySell bool) (*tradeapi.GetPortfolioResult, error)
	// Получение канала orderBooksChan
	GetOrderBooksChan() chan *tradeapi.OrderBookEvent
	// Получение канала orderTradeChan
	GetOrderTradeChan() chan *tradeapi.TradeEvent
	// Получение канала orderChan
	GetOrderChan() chan *tradeapi.OrderEvent
	// Получение канала ошибок
	GetErrorChan() chan error
	// Закрытие подключения
	CloseConnection()
}

// FinamClient
type FinamClient struct {
	token          string
	clientId       string
	ctx            context.Context
	connection     *grpc.ClientConn
	portfolio      tradeapi.PortfoliosClient
	orders         tradeapi.OrdersClient
	stops          tradeapi.StopsClient
	events         tradeapi.EventsClient
	orderBooksChan chan *tradeapi.OrderBookEvent
	orderTradeChan chan *tradeapi.TradeEvent
	orderChan      chan *tradeapi.OrderEvent
	errChan        chan error
}

func NewFinamClient(clientId, token string, ctx context.Context) (IFinamClient, error) {
	endpoint := "trade-api.finam.ru:443"

	tlsConfig := tls.Config{MinVersion: tls.VersionTLS12}

	conn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(credentials.NewTLS(&tlsConfig)))
	if err != nil {
		return nil, err
	}

	client := &FinamClient{
		token:          token,
		clientId:       clientId,
		ctx:            ctx,
		connection:     conn,
		portfolio:      tradeapi.NewPortfoliosClient(conn),
		orders:         tradeapi.NewOrdersClient(conn),
		stops:          tradeapi.NewStopsClient(conn),
		events:         tradeapi.NewEventsClient(conn),
		orderBooksChan: make(chan *tradeapi.OrderBookEvent),
		orderTradeChan: make(chan *tradeapi.TradeEvent),
		orderChan:      make(chan *tradeapi.OrderEvent),
		errChan:        make(chan error),
	}
	return client, err
}

func (f *FinamClient) CloseConnection() {
	f.connection.Close()
	close(f.orderBooksChan)
	close(f.orderTradeChan)
	close(f.orderChan)
	close(f.errChan)
}

func (f *FinamClient) CreateRequestContext() {
	f.ctx = grpcMetadata.AppendToOutgoingContext(f.ctx, "x-api-key", f.token)
}
