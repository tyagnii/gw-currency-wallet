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
	UUID, err := generateWalletID()
	if err != nil {
		return err
	}

	_, err = p.PGConn.Exec(p.ctx,
		`INSERT INTO users(name,email,wallet_id)  values($1,$2,$3)`,
		user.Username, user.Email, UUID)
	if err != nil {
		return err
	}

	return nil
}

func (p *PGConnector) GetUser(ctx context.Context, u models.User) (models.User, error) {
	r, err := p.PGConn.Query(ctx, "SELECT * FROM users WHERE username = $1", u.Username)
	if err != nil {
		return models.User{}, err
	}

	err = r.Scan(&u.ID, &u.Email, &u.Username)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}

func (p *PGConnector) CreateWallet(ctx context.Context, w models.Wallet) error {
	_, err := p.PGConn.Exec(ctx,
		"INSERT INTO wallets(uuid, balanceRUB, balanceUSD, balanceEUR) values ($1,$2,$3,$4)",
		w.UUID, 0, 0, 0)
	if err != nil {
		return err
	}

	return nil
}

func (p *PGConnector) Deposit(ctx context.Context, w models.Wallet) error {
	_, err := p.PGConn.Exec(ctx,
		"UPDATE wallets SET "+
			"balanceRUB= balanceRUB + $2, balanceUSD= balanceUSD + $3, balanceEUR= balanceEUR + $4)"+
			"WHERE uuid = $1",
		w.UUID, w.Balance.RUB, w.Balance.USD, w.Balance.EUR)
	if err != nil {
		return err
	}

	return nil
}

func (p *PGConnector) Withdraw(ctx context.Context, w models.Wallet) error {
	_, err := p.PGConn.Exec(ctx,
		"UPDATE wallets SET "+
			"balanceRUB= balanceRUB - $2, balanceUSD= balanceUSD - $3, balanceEUR= balanceEUR - $4)"+
			"WHERE uuid = $1",
		w.UUID, w.Balance.RUB, w.Balance.USD, w.Balance.EUR)
	if err != nil {
		return err
	}

	return nil
}
