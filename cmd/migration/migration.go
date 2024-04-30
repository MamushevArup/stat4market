package migration

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/pressly/goose/v3"
	"log"
	"os"
	"time"
)

const (
	defaultAttempts = 20
	defaultTimeout  = time.Second
	dir             = "migration"
	driverName      = "clickhouse"
)

var (
	NoMigrationFiles = errors.New("goose: no migration files")
	NoDbConnection   = errors.New("goose: environment variable not declared: CLICKHOUSE_URL")
)

func Run() error {
	databaseURL, ok := os.LookupEnv("CLICKHOUSE_URL")
	if !ok || len(databaseURL) == 0 {
		return NoDbConnection
	}

	var (
		attempts = defaultAttempts
		err      error
		db       *sql.DB
	)

	err = goose.SetDialect(driverName)
	if err != nil {
		return fmt.Errorf("goose: set dialect error: %w", err)
	}

	for attempts > 0 {
		db, err = sql.Open(driverName, databaseURL)
		if err == nil {
			break
		}

		log.Printf("goose: clickhouse db is trying to connect, attempts left: %d", attempts)
		time.Sleep(defaultTimeout)
		attempts--
	}

	if err != nil {
		return fmt.Errorf("goose: clickhouse db connect error: %w", err)
	}

	defer func() { _ = db.Close() }()

	err = goose.Up(db, dir)
	if err != nil {
		if errors.Is(err, goose.ErrAlreadyApplied) {
			log.Println("goose: up migration applied")
			return nil
		}
		if errors.Is(err, goose.ErrNoMigrationFiles) {
			return NoMigrationFiles
		}
		return err
	}
	return nil
}
