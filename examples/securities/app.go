package securitiesexample

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

	res, err := client.GetSecurities()
	if err != nil {
		log.Panicln(err)
	}

	log.Println(res.Securities)
}
