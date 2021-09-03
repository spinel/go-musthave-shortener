package web

import (
	"errors"

	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/pkg"
)

// EntityRepo is a repo for objects stored in memory(map).
type EntityRepo struct {
	memory map[string]model.Entity
}

// NewEntityRepo is a builder of repository.
func NewEntityRepo(db map[string]model.Entity) *EntityRepo {
	var repo EntityRepo
	repo.memory = db

	return &repo
}

// GetEntityBy - retrive entity by id.
func (repo *EntityRepo) GetEntityBy(id string) (*model.Entity, error) {
	if entity, ok := repo.memory[id]; ok {
		return &entity, nil
	}
	return nil, errors.New("not found")
}

// SaveEntity - saves given model by id.
func (repo *EntityRepo) SaveEntity(id string, entity model.Entity) error {
	repo.memory[id] = entity
	return nil
}

// IncludesCode check if id exists in repo.
func (repo *EntityRepo) IncludesCode(id string) bool {
	_, found := repo.memory[id]
	return found
}

func (repo *EntityRepo) GetCode(url string) (string, error) {
	if len(url) < 1 {
		return "", errors.New("wrong url")
	}
	var code string
	var err error
	for {
		code, err = pkg.NewGeneratedString()
		if err != nil {
			return "", err
		}

		if !repo.IncludesCode(string(code)) {
			break
		}
	}
	entity := model.Entity{
		URL: url,
	}

	err = repo.SaveEntity(code, entity)
	if err != nil {
		return "", err
	}
	return code, nil
}
