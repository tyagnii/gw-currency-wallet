package handlers

import (
	"context"
	"github.com/tyagnii/gw-currency-wallet/internal/db"
	exchanger_client "github.com/tyagnii/gw-currency-wallet/pkg/exchanger-client"
	"go.uber.org/zap"
)

type Handler struct {
	dbconn  db.DBConnector
	eClient exchanger_client.ExchangerClient
	sLogger *zap.SugaredLogger
}

func NewHandler(c context.Context, sLogger *zap.SugaredLogger) (*Handler, error) {
	pg, err := db.NewPGConnector(c, "")
	if err != nil {
		return nil, err
	}
	return &Handler{dbconn: pg, sLogger: sLogger}, nil
}
