package handlers

import "github.com/tyagnii/gw-currency-wallet/internal/db"

type Handler struct {
	dbconn db.DBConnector
}
