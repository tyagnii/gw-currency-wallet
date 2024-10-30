package db

import (
	"context"
)

type DBConnector interface {
	CreateWallet(ctx context.Context) (id string, err error)
	GetWalletBalance(ctx context.Context, uuid string) (balance int, err error)
	ChangeWalletBalance(ctx context.Context, uuid string, op string, amount int) (balance int, err error)
}
