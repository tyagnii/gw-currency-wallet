package db

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"os"
	"strings"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const migrationsPath string = "file://internal/migrations"

func buildConnString() string {
	builder := strings.Builder{}
	builder.WriteString("postgres://")
	builder.WriteString(os.Getenv("PGUSER"))
	builder.WriteString(":")
	builder.WriteString(os.Getenv("PGPASSWORD"))
	builder.WriteString("@")
	builder.WriteString(os.Getenv("PGHOST"))
	builder.WriteString(":")
	builder.WriteString(os.Getenv("PGPORT"))
	builder.WriteString("/")
	builder.WriteString(os.Getenv("PGDATABASE"))
	builder.WriteString("?sslmode=")
	builder.WriteString(os.Getenv("PGSSLMODE"))

	return builder.String()
}

func InitSchema() {
	// TODO: refactor connection to DB
	var m *migrate.Migrate
	var err error

	connectionString := buildConnString()

	for {
		m, err = migrate.New(migrationsPath,
			connectionString)
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
