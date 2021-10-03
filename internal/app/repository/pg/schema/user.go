package schema

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
)

type Entity struct {
	ID        int        `pg:"id,notnull,pk"`
	Code      string     `pg:"code"`
	URL       string     `pg:"url"`
	UserUUID  uuid.UUID  `pg:"user_uuid"`
	CreatedAt time.Time  `pg:"created_at,notnull"`
	UpdatedAt time.Time  `pg:"updated_at,notnull"`
	DeletedAt *time.Time `pg:"deleted_at"`
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

// NewEntityFromCanonical creates a new Entity DB object from canonical model.
func NewEntityFromCanonical(obj model.Entity) Entity {
	return Entity{
		ID:        obj.ID,
		Code:      obj.Code,
		URL:       obj.URL,
		UserUUID:  obj.UserUUID,
		CreatedAt: obj.CreatedAt,
		UpdatedAt: obj.UpdatedAt,
		DeletedAt: obj.DeletedAt,
	}
}

// ToCanonical converts a DB object to canonical model.
func (e Entity) ToCanonical() (model.Entity, error) {
	return model.Entity{
		ID:        e.ID,
		Code:      e.Code,
		URL:       e.URL,
		UserUUID:  e.UserUUID,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: e.DeletedAt,
	}, nil
}
