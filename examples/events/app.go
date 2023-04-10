package main

import (
	"context"
	"log"

	finamclient "github.com/DBoyara/FinamTradeGo"
	"github.com/DBoyara/FinamTradeGo/tradeapi"
)

func main() {
	ctx := context.Background()
	client, err := finamclient.NewFinamClient("clientId", "token", ctx)
	if err != nil {
		log.Panicln(err)
	}

	defer client.CloseConnection()

	in := &tradeapi.OrderBookSubscribeRequest{
		RequestId:     "32ef5786-e887",
		SecurityCode:  "GAZP",
		SecurityBoard: "TQBR",
	}

	go client.SubscribeOrderBook(in)

	chan1 := client.GetOrderBooksChan()
	chan_err := client.GetErrorChan()

	for {
		select {
		case res := <-chan1:
			log.Println("Response from orderBooksChan", res)
		case err := <-chan_err:
			log.Println("Response from errorChan", err)
			break
		}
	}

}
