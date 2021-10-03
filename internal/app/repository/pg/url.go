package pg

import (
	"context"
	"fmt"

	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/repository/pg/schema"
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
func (urlRepo *URLPgRepo) CreateURL(ctx context.Context, obj *model.Entity) (*model.Entity, error) {
	dbObj := schema.NewEntityFromCanonical(*obj)
	result, err := urlRepo.db.Model(&dbObj).
		OnConflict("(url) DO NOTHING").
		Insert()

	if err != nil {
		return nil, err
	}

	if result.RowsAffected() == 0 {
		existEntity, err := urlRepo.getByURL(obj.URL)
		if err != nil {
			return nil, err
		}

		obj, err := existEntity.ToCanonical()
		if err != nil {
			return &model.Entity{}, fmt.Errorf("conveting to canonical model: %w", err)
		}

		return &obj, nil
	}

	return nil, nil
}

// GetURL retrives entity by code.
func (urlRepo *URLPgRepo) GetURL(ctx context.Context, urlCode string) (*model.Entity, error) {
	dbObj := schema.Entity{}

	err := urlRepo.db.Model(&dbObj).
		Where("code = ?", urlCode).
		Where(notDeleted).
		Select()

	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() {
			return nil, nil
		}
		return nil, err
	}

	obj, err := dbObj.ToCanonical()
	if err != nil {
		return nil, fmt.Errorf("conveting to canonical model: %w", err)
	}

	return &obj, nil
}

// GetByUser retrive entities by user UUID.
func (urlRepo *URLPgRepo) GetByUser(ctx context.Context, userUUID uuid.UUID) ([]model.Entity, error) {
	var dbObjPool []schema.Entity

	err := urlRepo.db.Model(&dbObjPool).
		Where("user_uuid = ?", userUUID.String()).
		Where(notDeleted).
		Select()

	if err != nil {
		return nil, err
	}

	var objPool []model.Entity
	for _, dbObj := range dbObjPool {
		obj, err := dbObj.ToCanonical()
		if err != nil {
			return nil, fmt.Errorf("conveting to canonical model: %w", err)
		}

		objPool = append(objPool, obj)
	}

	return objPool, nil
}

// SaveBatch uses to save array of entities.
func (urlRepo *URLPgRepo) SaveBatch(ctx context.Context, objPool []model.Entity) error {
	var dbObjPool []schema.Entity
	for _, obj := range objPool {
		dbObj := schema.NewEntityFromCanonical(obj)

		dbObjPool = append(dbObjPool, dbObj)
	}

	tx, err := urlRepo.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// prepare each object to commit.
	for _, dbObj := range dbObjPool {
		_, err := tx.Model(&dbObj).
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
func (urlRepo *URLPgRepo) getByURL(url string) (*schema.Entity, error) {
	dbObj := &schema.Entity{}

	err := urlRepo.db.Model(dbObj).
		Where("url = ?", url).
		Where(notDeleted).
		Select()

	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() { //not found
			return nil, nil
		}
		return nil, err
	}

	return dbObj, nil
}
