package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
)

type PGConnector struct {
	PGConn *pgx.Conn
	ctx    context.Context
}

func NewPGConnector(ctx context.Context, connectionString string) (*PGConnector, error) {
	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}
	return &PGConnector{PGConn: conn, ctx: ctx}, nil
}

func (p *PGConnector) CreateUser(ctx context.Context, user models.User) error {
	return nil
}

func (p *PGConnector) GetUser(ctx context.Context, u models.User) (models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PGConnector) CreateWallet(ctx context.Context, w models.Wallet) error {
	//TODO implement me
	panic("implement me")
}

func (p *PGConnector) Deposit(ctx context.Context, w models.Wallet) error {
	//TODO implement me
	panic("implement me")
}

func (p *PGConnector) Withdraw(ctx context.Context, w models.Wallet) error {
	//TODO implement me
	panic("implement me")
}
