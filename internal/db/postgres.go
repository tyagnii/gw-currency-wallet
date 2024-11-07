package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/tyagnii/gw-currency-wallet/internal/models"
	"reflect"
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

func (p *PGConnector) Exchange(ctx context.Context, w models.Wallet, req models.ExchangeReq) (models.Wallet, error) {
	// Experimental!
	// Implement reflection identification of fields in case
	// we have huge number of currencies
	rFromBalance := reflect.ValueOf(w.Balance).Elem().FieldByName(req.FromCurrency)
	rToBalance := reflect.ValueOf(w.Balance).Elem().FieldByName(req.ToCurrency)
	rRateTo := reflect.ValueOf(req.Rate).Elem().FieldByName(req.ToCurrency)
	rRateFrom := reflect.ValueOf(req.Rate).Elem().FieldByName(req.FromCurrency)

	if rFromBalance.IsValid() && rToBalance.IsValid() {
		rFromBalance.SetFloat(rFromBalance.Float() - req.Amount)
		rToBalance.SetFloat(rToBalance.Float() + req.Amount*rRateTo.Float()*rRateFrom.Float())
	}

	// Start transaction
	tx, err := p.PGConn.Begin(ctx)
	if err != nil {
		return models.Wallet{}, err
	}

	_, err = tx.Exec(
		ctx,
		`UPDATE wallets
			SET balanceRUB = $1,
			SET balanceUSD = $2,
			SET balanceEUR = $3
			WHERE uuid = $4;`,
		w.Balance.RUB,
		w.Balance.USD,
		w.Balance.EUR,
		w.UUID)
	if err != nil {
		err = tx.Rollback(ctx)
		if err != nil {
			err = fmt.Errorf("failed to rollback transaction: %w", err)
		}
		return models.Wallet{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		err = tx.Rollback(ctx)
		if err != nil {
			err = fmt.Errorf("failed to rollback transaction: %w", err)
		}
		return models.Wallet{}, err
	}

	return w, nil
}

func (p *PGConnector) CreateUser(ctx context.Context, user models.User) error {
	UUID, err := generateWalletID()
	if err != nil {
		return err
	}

	w := models.Wallet{UUID: UUID}

	if err = p.CreateWallet(ctx, w); err != nil {
		return err
	}

	_, err = p.PGConn.Exec(p.ctx,
		`INSERT INTO users(username, password, email, wallet_id) values($1,$2,$3,$4)`,
		user.Username, user.Password, user.Email, UUID)
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

func (p *PGConnector) GetWalletByUsername(ctx context.Context, username string) (models.Wallet, error) {
	var w models.Wallet
	r, err := p.PGConn.Query(ctx,
		`SELECT wallets.UUID, wallets.balanceRUB, wallets.balanceUSD, wallets.balanceEUR
			FROM wallets
			WHERE username = $1
			JOIN users
         	ON wallets.uuid = users.wallets_uuid`, username)
	if err != nil {
		return models.Wallet{}, err
	}

	err = r.Scan(&w.UUID, &w.Balance.RUB, &w.Balance.USD, &w.Balance.EUR)
	if err != nil {
		return models.Wallet{}, err
	}

	return w, nil
}

// GetBalance returns current wallet balance
func (p *PGConnector) GetBalance(ctx context.Context, u models.User) (models.Wallet, error) {
	var w models.Wallet

	r, err := p.PGConn.Query(ctx,
		`SELECT wallets.balanceRUB, wallets.balanceUSD, wallets.balanceEUR
		FROM wallets
		JOIN (SELECT * FROM users where users.username = $1) as t1
		    ON wallets.uuid = t1.wallet_id`, u.Username)
	if err != nil {
		return models.Wallet{}, err
	}

	r.Next()
	err = r.Scan(&w.Balance.RUB, &w.Balance.USD, &w.Balance.EUR)
	if err != nil {
		return models.Wallet{}, fmt.Errorf("erorr during scan rows to struct: %w", err)
	}

	return w, nil
}
