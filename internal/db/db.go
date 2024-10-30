package db

import (
	"context"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
)

type DBConnector interface {
	CreateUser(ctx context.Context, u models.User) error
	GetUser(ctx context.Context, u models.User) (models.User, error)
	CreateWallet(ctx context.Context, w models.Wallet) error
	Deposit(ctx context.Context, w models.Wallet) error
	Withdraw(ctx context.Context, w models.Wallet) error
}
