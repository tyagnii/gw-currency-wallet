package db

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

const migrationsPath string = "file://internal/db/migrations"

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

func InitSchema(sLogger *zap.SugaredLogger) error {
	// TODO: refactor connection to DB
	// 		channel for errors
	var m *migrate.Migrate
	var err error
	var timeout time.Duration

	timeout, err = time.ParseDuration(os.Getenv("MIGRATE_TIMEOUT"))
	if err != nil {
		return err
	}

	connectionString := buildConnString()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	go func(ctx context.Context) {
		for {
			m, err = migrate.New(migrationsPath, connectionString)
			if err != nil {
				sLogger.Errorf("Could not connect to database: %s", err)
			} else {
				break
			}
			switch ctx.Err() {
			case context.Canceled:
				break
			case context.DeadlineExceeded:
				break
			default:
				time.Sleep(5 * time.Second)
			}
		}
	}(ctx)

	<-ctx.Done()

	if m == nil {
		return fmt.Errorf("could not connect to database")
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
