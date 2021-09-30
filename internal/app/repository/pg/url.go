package pg

import (
	"context"

	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
)

type URLPgRepo struct {
	db *DB
}

// NewURLPgRepo is a URLPgRepo builder.
func NewURLPgRepo(db *DB) *URLPgRepo {
	return &URLPgRepo{db: db}
}

// Ping checks db connection.
func (urlRepo *URLPgRepo) Ping() bool {
	_, err := urlRepo.db.Exec("SELECT 1")
	return err == nil
}

// CreateURL save entity to db.
func (urlRepo *URLPgRepo) CreateURL(entity *model.Entity) (*model.Entity, error) {
	result, err := urlRepo.db.Model(entity).
		OnConflict("(url) DO NOTHING").
		Insert()

	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		existEntity, err := urlRepo.getByURL(entity.URL)
		if err != nil {
			return nil, err
		}

		return existEntity, nil
	}

	return nil, nil
}

// GetURL retrives entity by code.
func (urlRepo *URLPgRepo) GetURL(urlCode string) (*model.Entity, error) {
	entity := &model.Entity{}
	err := urlRepo.db.Model(entity).
		Where("code = ?", urlCode).
		Where(notDeleted).
		Select()
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() { //not found
			return nil, nil
		}
		return nil, err
	}
	return entity, nil
}

// GetByUser retrive entities by user UUID.
func (urlRepo *URLPgRepo) GetByUser(ctx context.Context, userUUID uuid.UUID) []model.Entity {
	var entityPool []model.Entity
	err := urlRepo.db.Model(&entityPool).
		Where("user_uuid = ?", userUUID.String()).
		Where(notDeleted).
		Select()
	if err != nil {
		return nil
	}

	return entityPool
}

// SaveBatch uses to save array of entities.
func (urlRepo *URLPgRepo) SaveBatch(ctx context.Context, entities []model.Entity) error {
	tx, err := urlRepo.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Close()

	// prepare each entities to commit.
	for _, entity := range entities {
		_, err := tx.Model(&entity).
			Insert()
		if err != nil {
			return err
		}
	}

	// commit on success.
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// getByURL retrives entity by url.
func (urlRepo *URLPgRepo) getByURL(url string) (*model.Entity, error) {
	entity := &model.Entity{}
	err := urlRepo.db.Model(entity).
		Where("url = ?", url).
		Where(notDeleted).
		Select()
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() { //not found
			return nil, nil
		}
		return nil, err
	}
	return entity, nil
}
