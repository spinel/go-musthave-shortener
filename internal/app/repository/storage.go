package repository

import (
	"github.com/pkg/errors"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/repository/pg"
)

type Storage struct {
	Pg       *pg.DB
	EntityPg UrlStorer
}

// NewStorage is a gob storage builder
func NewStorage(cfg *config.Config) (*Storage, error) {
	pgDB, err := pg.Dial(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "pgdb.Dial failed")
	}

	entityRepoPg := pg.NewUrlPgRepo(pgDB)

	s := &Storage{
		Pg:       pgDB,
		EntityPg: entityRepoPg,
	}

	return s, nil
}
