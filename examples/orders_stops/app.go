package ordersstops

import (
	"context"
	"log"

	finamclient "github.com/evsamsonov/FinamTradeGo/v2"
	"github.com/evsamsonov/FinamTradeGo/v2/tradeapi"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	ctx := context.Background()
	client, err := finamclient.NewFinamClient("ClientId", "token", ctx)
	if err != nil {
		log.Panicln(err)
	}

	defer client.CloseConnection()

	in := &tradeapi.NewOrderRequest{
		ClientId:      "ClientId",
		SecurityBoard: "FUT",
		SecurityCode:  "SiM3",
		BuySell:       2,
		Quantity:      1,
		UseCredit:     false,
		Price:         &wrapperspb.DoubleValue{},
		Property:      1,
		Condition:     &tradeapi.OrderCondition{},
		ValidBefore:   &tradeapi.OrderValidBefore{},
	}

	res_order, err := client.NewOrder(in)
	if err != nil {
		log.Panicln(err)
	}

	log.Printf("Id заявки: %d", res_order.TransactionId)

	in_stop := &tradeapi.NewStopRequest{
		ClientId:      "ClientId",
		SecurityBoard: "FUT",
		SecurityCode:  "SiM3",
		BuySell:       1,
		StopLoss: &tradeapi.StopLoss{
			ActivationPrice: 69000,
			Price:           69000,
			MarketPrice:     true,
			Quantity: &tradeapi.StopQuantity{
				Value: 1,
				Units: tradeapi.StopQuantityUnits_STOP_QUANTITY_UNITS_LOTS,
			},
			UseCredit: false,
		},
		TakeProfit: &tradeapi.TakeProfit{
			ActivationPrice: 70000,
			CorrectionPrice: &tradeapi.StopPrice{},
			SpreadPrice:     &tradeapi.StopPrice{},
			MarketPrice:     true,
			Quantity: &tradeapi.StopQuantity{
				Value: 1,
				Units: tradeapi.StopQuantityUnits_STOP_QUANTITY_UNITS_LOTS,
			},
			UseCredit: false,
		},
		ExpirationDate: &timestamppb.Timestamp{},
		LinkOrder:      int64(res_order.TransactionId),
		ValidBefore:    &tradeapi.OrderValidBefore{},
	}

	res_stop, err := client.NewStop(in_stop)
	if err != nil {
		log.Panicln(err)
	}

	log.Printf("Id стопа: %d", res_stop.GetStopId())
}
