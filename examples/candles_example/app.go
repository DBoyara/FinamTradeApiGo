package candlesexample

import (
	"context"
	"log"

	finamclient "github.com/DBoyara/FinamTradeGo"
	"github.com/DBoyara/FinamTradeGo/tradeapi"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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
