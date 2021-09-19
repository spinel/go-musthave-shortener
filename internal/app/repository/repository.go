package repository

import (
	"context"

	"github.com/spinel/go-musthave-shortener/internal/app/model"
)

// URLShortener defines ntity operations.
type URLShortener interface {
	GetEntityBy(id string) (*model.Entity, error)
	SaveEntity(id string, entity model.Entity) error
	GetCode(ctx context.Context, url string) (string, error)
	GetByUser(ctx context.Context) []model.URLMapping
	GetMemory() model.MemoryMap
}
