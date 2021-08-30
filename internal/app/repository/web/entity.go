package web

import (
	"github.com/spinel/go-musthave-shortener-tpl/internal/app/model"
)

// EntityRepo is a repo for objects.
type EntityRepo struct {
	Memory map[string]model.Entity
}

// NewEntityRepo is a builder of repository.
func NewEntityRepo(db map[string]model.Entity) *EntityRepo {
	var repo EntityRepo
	repo.Memory = db

	return &repo
}

// GetEntityBy - retrive entity by id.
func (repo *EntityRepo) GetEntityBy(id string) (model.Entity, error) {
	entity := repo.Memory[id]
	return entity, nil
}

// SaveEntity - saves given model by id.
func (repo *EntityRepo) SaveEntity(id string, entity model.Entity) error {
	repo.Memory[id] = entity
	return nil
}

// IncludesCode check if id exists in repo.
func (repo *EntityRepo) IncludesCode(id string) bool {
	_, result := repo.Memory[id]
	return result
}
