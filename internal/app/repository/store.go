package repository

import (
	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/repository/web"
)

// Store main struct.
type Store struct {
	memoryDB map[string]model.Entity
	Entity   Repositorer
}

// New store builder.
func New() (*Store, error) {
	db := make(map[string]model.Entity)
	store := Store{
		memoryDB: db,
		Entity:   web.NewEntityRepo(db),
	}

	return &store, nil
}
