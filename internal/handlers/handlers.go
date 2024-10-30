package handlers

import (
	"context"
	"github.com/tyagnii/gw-currency-wallet/internal/db"
)

type Handler struct {
	dbconn db.DBConnector
}

func NewHandler(c context.Context) (*Handler, error) {
	pg, err := db.NewPGConnector(c, "")
	if err != nil {
		return nil, err
	}
	return &Handler{dbconn: pg}, nil
}
