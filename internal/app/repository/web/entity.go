package web

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/pkg"
)

// EntityRepo is a repo for objects stored in memory(map).
type EntityRepo struct {
	memory model.MemoryMap
}

// NewEntityRepo is a builder of repository.
func NewEntityRepo(db model.MemoryMap) *EntityRepo {
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
func (repo *EntityRepo) includesCode(id string) bool {
	_, found := repo.memory[id]
	return found
}

func (repo *EntityRepo) GetCode(ctx context.Context, url string) (string, error) {
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

		if !repo.includesCode(string(code)) {
			break
		}
	}

	userUUIDString := ctx.Value(model.CookieContextName).(string)
	userUUID, _ := uuid.Parse(userUUIDString)

	entity := model.Entity{
		URL:      url,
		UserUUID: userUUID,
	}

	err = repo.SaveEntity(code, entity)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (repo *EntityRepo) GetByUser(ctx context.Context, cfg *config.Config) []model.URLMapping {
	userUUIDString := ctx.Value(model.CookieContextName).(string)
	userUUID, _ := uuid.Parse(userUUIDString)

	var urlMappingPool []model.URLMapping

	for code, entity := range repo.memory {
		if entity.UserUUID == userUUID {
			box := model.URLMapping{
				ShortURL:    fmt.Sprintf("%s/%s", cfg.BaseURL, code),
				OriginalURL: entity.URL,
			}
			urlMappingPool = append(urlMappingPool, box)
		}
	}
	return urlMappingPool
}

func (repo *EntityRepo) GetMemory() model.MemoryMap {
	return repo.memory
}

func (repo *EntityRepo) Ping() bool {
	_, ok := repo.memory["testtest"]
	return ok
}

func (repo *EntityRepo) SaveBatch(ctx context.Context, batch []model.RequestBatchURLS) error {
	return nil
}
