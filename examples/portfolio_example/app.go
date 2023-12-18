package portfolioexample

import (
	"context"
	"log"

	finamclient "github.com/evsamsonov/FinamTradeGo/v2"
)

func main() {
	ctx := context.Background()
	client, err := finamclient.NewFinamClient("ClientId", "token", ctx)
	if err != nil {
		log.Panicln(err)
	}

	defer client.CloseConnection()

	res, err := client.GetPortfolio(true, true, true, true)
	if err != nil {
		log.Panicln(err)
	}

	log.Printf("Входящая оценка портфеля в рублях: %f", res.Balance)
	log.Println(res.Positions)
}
