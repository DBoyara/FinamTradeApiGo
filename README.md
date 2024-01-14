# FinamTradeGo

Finam GRPC-client on Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/DBoyara/FinamTradeGo)](https://goreportcard.com/report/github.com/DBoyara/FinamTradeGo) [![GoDoc](https://godoc.org/github.com/DBoyara/FinamTradeGo?status.svg)](https://pkg.go.dev/github.com/DBoyara/FinamTradeGo)
[![Maintainability](https://api.codeclimate.com/v1/badges/d11095947b38e5085c8f/maintainability)](https://codeclimate.com/github/DBoyara/FinamTradeApiGo/maintainability)



## Installation


```bash
  go get github.com/evsamsonov/FinamTradeGo/v2 
```



[portfolios.proto](..%2Ftrade-api-docs%2Fcontracts%2Fproto%2Ftradeapi%2Fv1%2Fportfolios.proto)

## Examples

### Пример получения свечей
```go
func main() {
	ctx := context.Background()
	client, err := finamclient.NewFinamClient("ClientId", "token", ctx)
	if err != nil {
		log.Panicln(err)
	}

	defer client.CloseConnection()

	in := &tradeapi.GetDayCandlesRequest{
		SecurityBoard: "TQBR",
		SecurityCode:  "SBER",
		TimeFrame:     1,
		Interval: &tradeapi.DayCandleInterval{
			From: &date.Date{
				Year:  2023,
				Month: 6,
				Day:   5,
			},
			To: &date.Date{
				Year:  2023,
				Month: 6,
				Day:   9,
			},
		},
	}
	res, err := client.GetDayCandles(in)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res.GetCandles())

	response := &tradeapi.GetIntradayCandlesRequest{
		SecurityBoard: "TQBR",
		SecurityCode:  "SBER",
		TimeFrame:     2,
		Interval: &tradeapi.IntradayCandleInterval{
			From: &timestamppb.Timestamp{
				Seconds: 1686286800,
				Nanos:   0,
			},
			Count: 10,
		},
	}
	res2, err := client.GetIntradayCandles(response)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res2.GetCandles())

}
```

### Пример выставления заявки на покупку по рынку. Не рекомендация к действию.
```go
func main() {
	ctx := context.Background()
	client, err := NewFinamClient("clientId", "token", ctx)
	if err != nil {
		log.Panicln(err)
	}

	defer client.CloseConnection()

	in := &ta.NewOrderRequest{
		ClientId:      "ClientId",
		SecurityBoard: "FUT",
		SecurityCode:  "SiM3",
		BuySell:       2,
		Quantity:      1,
		UseCredit:     false,
		Price:         &wrapperspb.DoubleValue{},
		Property:      1,
		Condition:     &ta.OrderCondition{},
		ValidBefore:   &ta.OrderValidBefore{},
	}

	res_order, err := client.NewOrder(in)
	if err != nil {
		log.Panicln(err)
	}

	log.Printf("Id заявки: %d", res_order.TransactionId)
}
```

### Пример выставления стопов по рынку. Не рекомендация к действию. Подробнее про стоп-завявки [тут](https://finamweb.github.io/trade-api-docs/usage#%D1%81%D1%82%D0%BE%D0%BF-%D0%B7%D0%B0%D1%8F%D0%B2%D0%BA%D0%B8)
```go
func main() {
	...

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
			SpreadPrice: &tradeapi.StopPrice{},
			MarketPrice: true,
			Quantity: &tradeapi.StopQuantity{
				Value: 1,
				Units: tradeapi.StopQuantityUnits_STOP_QUANTITY_UNITS_LOTS,
			},
			UseCredit: false,
		},
		ExpirationDate: &timestamppb.Timestamp{},
		LinkOrder: int64(res_order.TransactionId),
		ValidBefore: &tradeapi.OrderValidBefore{},
	}

	res_stop, err := client.NewStop(in_stop)
	if err != nil {
		log.Panicln(err)
	}

	log.Printf("Id стопа: %d", res_stop.GetStopId())
}
```

### Пример подписки на биржевой стакан. Не рекомендация к действию.
```go
func main() {
	ctx := context.Background()
	client, err := NewFinamClient("clientId", "token", ctx)
	if err != nil {
		log.Panicln(err)
	}

	defer client.CloseConnection()

	in := &tradeapi.OrderBookSubscribeRequest{
		RequestId:     "ffc38cb7-2072",
		SecurityCode:  "GAZP",
		SecurityBoard: "TQBR",
	}

	go client.SubscribeOrderBook(in)

	chan1 := client.GetOrderBooksChan()
	chan_err := client.GetErrorChan()

	for {
		select {
		case res := <-chan1:
			fmt.Println("Response from orderBooksChan", res)
		case err := <-chan_err:
			fmt.Println("Response from errorChan", err)
			client.CloseConnection()
			break
		}
	}
}
```

## Authors

- [@DBoyara](https://www.github.com/DBoyara)s


## License

ДАННОЕ ПРОГРАММНОЕ ОБЕСПЕЧЕНИЕ ПРЕДОСТАВЛЯЕТСЯ «КАК ЕСТЬ», БЕЗ КАКИХ-ЛИБО ГАРАНТИЙ, ЯВНО ВЫРАЖЕННЫХ ИЛИ ПОДРАЗУМЕВАЕМЫХ, ВКЛЮЧАЯ ГАРАНТИИ ТОВАРНОЙ ПРИГОДНОСТИ, СООТВЕТСТВИЯ ПО ЕГО КОНКРЕТНОМУ НАЗНАЧЕНИЮ И ОТСУТСТВИЯ НАРУШЕНИЙ, НО НЕ ОГРАНИЧИВАЯСЬ ИМИ. НИ В КАКОМ СЛУЧАЕ АВТОРЫ ИЛИ ПРАВООБЛАДАТЕЛИ НЕ НЕСУТ ОТВЕТСТВЕННОСТИ ПО КАКИМ-ЛИБО ИСКАМ, ЗА УЩЕРБ ИЛИ ПО ИНЫМ ТРЕБОВАНИЯМ, В ТОМ ЧИСЛЕ, ПРИ ДЕЙСТВИИ КОНТРАКТА, ДЕЛИКТЕ ИЛИ ИНОЙ СИТУАЦИИ, ВОЗНИКШИМ ИЗ-ЗА ИСПОЛЬЗОВАНИЯ ПРОГРАММНОГО ОБЕСПЕЧЕНИЯ ИЛИ ИНЫХ ДЕЙСТВИЙ С ПРОГРАММНЫМ ОБЕСПЕЧЕНИЕМ.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

The Finam clieint is open-sourced software licensed under the [GNU GPLv3](https://choosealicense.com/licenses/gpl-3.0/)
