package repository

import (
	"log"

	"github.com/pkg/errors"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/repository/pg"
)

type Storage struct {
	Pg       *pg.DB
	EntityPg URLStorer
}

// NewStorage is a gob storage builder
func NewStorage(cfg config.Config) (*Storage, error) {
	pgDB, err := pg.Dial(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "pgdb.Dial failed")
	}

	// Run Postgres migrations
	if pgDB != nil {
		log.Println("Running PostgreSQL migrations...")
		if err := runPgMigrations(cfg); err != nil {
			return nil, errors.Wrap(err, "runPgMigrations failed")
		}
	}

	entityRepoPg := pg.NewURLPgRepo(cfg, pgDB)

	s := &Storage{
		Pg:       pgDB,
		EntityPg: entityRepoPg,
	}

	return s, nil
}
