package model

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
)

// Entity is stored URL object.
// TODO: use net/url enstead of string.
// Current struct used as value of memory map.
type Entity struct {
	ID        int        `json:"id" pg:"id,notnull,pk"`
	Code      string     `json:"code" pg:"code"`
	URL       string     `json:"url" pg:"url"`
	UserUUID  uuid.UUID  `json:"user_uuid" pg:"user_uuid"`
	CreatedAt time.Time  `json:"created_at,omitempty" pg:"created_at,notnull"`
	UpdatedAt time.Time  `json:"updated_at,omitempty" pg:"updated_at,notnull"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" pg:"deleted_at"`
}

// BeforeInsert hooks into insert operations,
// setting createdAt and updatedAt to current time.
func (e *Entity) BeforeInsert(ctx context.Context) (context.Context, error) {
	now := time.Now()
	if e.CreatedAt.IsZero() {
		e.CreatedAt = now
	}
	if e.UpdatedAt.IsZero() {
		e.UpdatedAt = now
	}
	return ctx, nil
}

// ToURLMapping convert Entity to URLMapping struct.
func (e *Entity) ToURLMapping(cfg *config.Config) URLMapping {
	return URLMapping{
		OriginalURL: e.URL,
		ShortURL:    fmt.Sprintf("%s/%s", cfg.BaseURL, e.Code),
	}
}

// Result is response url struct for API.
type Result struct {
	URL string `json:"result"`
}

type MemoryMap map[string]Entity

type URLMapping struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}

type RequestBatchURLS struct {
	CorrelationID string `json:"correlation_id"`
	OriginalURL   string `json:"original_url"`
	ShortURL      string `json:"short_url"`
}
