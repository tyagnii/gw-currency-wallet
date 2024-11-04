package exchanger_client

import (
	"context"
	exchanger "github.com/tyagnii/gw-exchanger/gen/exchanger/v1"
)

type ExchangerClient struct {
	exchanger.ExchangeServiceClient
	// todo: add server connection string
}

func NewExchangerClient() *ExchangerClient {
	//exchanger.ExchangeRateResponse{Rate: 1.34}
	return &ExchangerClient{}
}

// GetExchangeRates calls to  Exchanger server and returns currency rates
func (e *ExchangerClient) GetExchangeRates(
	ctx context.Context,
	in *exchanger.Empty,
) (*exchanger.ExchangeRatesResponse, error) {

	return nil, nil

}

// GetExchangerRateForCurrency returns rates for provided currencies
func (e *ExchangerClient) GetExchangeRateForCurrency(
	ctx context.Context,
	in *exchanger.CurrencyRequest,
) (*exchanger.ExchangeRateResponse, error) {

	return nil, nil

}
