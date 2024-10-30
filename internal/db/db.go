package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type PGConnector struct {
	pgconn *pgx.Conn
	ctx    context.Context
}

func NewPGConnector(ctx context.Context, connectionString string) (*PGConnector, error) {
	conn, err := pgx.Connect(ctx, connectionString)
	if err != nil {
		return nil, err
	}
	return &PGConnector{pgconn: conn, ctx: ctx}, nil
}
