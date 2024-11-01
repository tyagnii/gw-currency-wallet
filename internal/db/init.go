package db

import (
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func init() {
	m, err := migrate.New("file://internal/db/migrations",
		"postgres://postgres:password@db:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = m.Migrate(1)
	if err != nil {
		panic(err)
	}
}
