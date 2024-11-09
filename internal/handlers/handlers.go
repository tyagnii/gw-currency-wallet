package handlers

import (
	"context"
	"github.com/tyagnii/gw-currency-wallet/internal/db"
	"github.com/tyagnii/gw-currency-wallet/internal/db/postgres"
	"github.com/tyagnii/gw-currency-wallet/pkg/cache"

	exchanger_client "github.com/tyagnii/gw-currency-wallet/pkg/exchanger-client"
	"go.uber.org/zap"
)

type Handler struct {
	dbconn  db.DBConnector
	eClient exchanger_client.ExchangerClient
	sLogger *zap.SugaredLogger
	*cache.Cache
}

func NewHandler(c context.Context, sLogger *zap.SugaredLogger) (*Handler, error) {
	pg, err := postgres.NewPGConnector(c, "")
	if err != nil {
		return nil, err
	}
	return &Handler{dbconn: pg, sLogger: sLogger}, nil
}
