package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func InitSchema() {
	// TODO: refactor connection to DB
	// 		read connection string from config
	var m *migrate.Migrate
	var err error
	for {
		m, err = migrate.New("file://internal/db/migrations",
			"postgres://postgres:password@db:5432/postgres?sslmode=disable")
		if err != nil {
			fmt.Println(err)
			time.Sleep(10 * time.Second)
		} else {
			break
		}
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
	}
}
