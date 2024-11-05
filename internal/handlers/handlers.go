package handlers

import (
	"context"
	"github.com/tyagnii/gw-currency-wallet/internal/db"
	exchanger_client "github.com/tyagnii/gw-currency-wallet/pkg/exchanger-client"
)

type Handler struct {
	dbconn  db.DBConnector
	eClient exchanger_client.ExchangerClient
}

func NewHandler(c context.Context) (*Handler, error) {
	pg, err := db.NewPGConnector(c, "")
	if err != nil {
		return nil, err
	}
	return &Handler{dbconn: pg}, nil
}
