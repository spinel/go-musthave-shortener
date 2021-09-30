package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
)

// UrlStorer defines Entity operations.
type URLStorer interface {
	Ping() bool
	GetURL(urlCode string) (*model.Entity, error)
	CreateURL(*model.Entity) (*model.Entity, error)
	GetByUser(ctx context.Context, userUUID uuid.UUID) []model.Entity
	SaveBatch(ctx context.Context, entities []model.Entity) error
}
