package exchanger_client

import (
	"context"
	exchanger "github.com/tyagnii/gw-exchanger/gen/exchanger/v1"
	"google.golang.org/grpc"
	"os"
)

type ExchangerClient struct {
	exchanger.ExchangeServiceClient
}

func NewExchangerClient() *ExchangerClient {
	//exchanger.ExchangeRateResponse{Rate: 1.34}
	serverAddress := os.Getenv("EXCHANGER_SERVER_ADDRESS")
	conn, err := grpc.NewClient(serverAddress)
	if err != nil {
		panic(err)
	}
	exClient := exchanger.NewExchangeServiceClient(conn)

	return &ExchangerClient{ExchangeServiceClient: exClient}
}

// GetExchangeRates calls to  Exchanger server and returns currency rates
func (e *ExchangerClient) GetExchangeRates(
	ctx context.Context,
	in *exchanger.Empty,
) (*exchanger.ExchangeRatesResponse, error) {
	return e.ExchangeServiceClient.GetExchangeRates(ctx, in)
}

// GetExchangeRateForCurrency returns rates for provided currencies
func (e *ExchangerClient) GetExchangeRateForCurrency(
	ctx context.Context,
	in *exchanger.CurrencyRequest,
) (*exchanger.ExchangeRateResponse, error) {
	return e.ExchangeServiceClient.GetExchangeRateForCurrency(ctx, in)
}
