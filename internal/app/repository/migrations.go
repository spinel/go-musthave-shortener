package repository

import (
	"github.com/golang-migrate/migrate"
	"github.com/pkg/errors"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
)

// runPgMigrations runs Postgres migrations
func runPgMigrations(cfg *config.Config) error {

	if cfg.PgMigrationsPath == "" {
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
