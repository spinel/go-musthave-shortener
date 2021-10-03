package pg

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
)

const (
	// Timeout is a Postgres timeout
	Timeout    = 5
	notDeleted = "deleted_at is null"
)

// DB is a shortcut structure to a Postgres DB
type DB struct {
	*pg.DB
}

// Dial creates new database connection to postgres
func Dial(cfg config.Config) (*DB, error) {
	pgOpts, err := pg.ParseURL(cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}

	pgDB := pg.Connect(pgOpts)

	_, err = pgDB.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}

	pgDB.WithTimeout(time.Second * time.Duration(Timeout))

	return &DB{pgDB}, nil
}
