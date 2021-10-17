package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
)

// UrlStorer defines Entity operations.
type URLStorer interface {
	Ping() bool
	GetURL(ctx context.Context, urlCode string) (*model.Entity, error)
	CreateURL(ctx context.Context, entity *model.Entity) (*model.Entity, error)
	GetByUser(ctx context.Context, userUUID uuid.UUID) ([]model.Entity, error)
	SaveBatch(ctx context.Context, entities []model.Entity) error
	DeleteBatch(ctx context.Context, entities []model.Entity) error
	EnqueueDelete(codes []string, userUUID uuid.UUID) error
}
