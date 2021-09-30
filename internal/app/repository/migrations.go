package repository

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"github.com/spinel/go-musthave-shortener/internal/app/config"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// runPgMigrations runs Postgres migrations
func runPgMigrations(cfg *config.Config) error {
	if cfg.PgMigrationsPath == "" {
		fmt.Println("ok\n")
		return nil
	}
	if cfg.DatabaseDSN == "" {
		return errors.New("No cfg.PgURL provided")
	}
	m, err := migrate.New(
		cfg.PgMigrationsPath,
		cfg.DatabaseDSN,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
