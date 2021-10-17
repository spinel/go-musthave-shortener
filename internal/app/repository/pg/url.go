package pg

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/model"

	"github.com/spinel/go-musthave-shortener/internal/app/repository/pg/schema"
)

type URLPgRepo struct {
	db               *DB
	WorkerDeleteChan chan model.BatchUserCode
}

// NewURLPgRepo is a URLPgRepo builder.
func NewURLPgRepo(cfg config.Config, db *DB) *URLPgRepo {
	workerDeleteChan := make(chan model.BatchUserCode)

	urlPgRepo := &URLPgRepo{
		db:               db,
		WorkerDeleteChan: workerDeleteChan,
	}
	urlPgRepo.batchDeleteWorker(cfg.BatchQueueSize)

	return urlPgRepo
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
		return errors.Wrap(err, "repo.SaveBatch")
	}
	defer tx.Rollback()

	// prepare each object to commit.
	for _, dbObj := range dbObjPool {
		_, err := tx.Model(&dbObj).
			Insert()
		if err != nil {
			return errors.Wrap(err, "repo.SaveBatch")
		}
	}

	// commit on success.
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "repo.SaveBatch")
	}

	return nil
}

// DeleteBatch uses to remove array of entities.
func (urlRepo *URLPgRepo) DeleteBatch(ctx context.Context, objPool []model.Entity) error {
	var dbObjPool []schema.Entity
	for _, obj := range objPool {
		dbObj := schema.NewEntityFromCanonical(obj)
		dbObjPool = append(dbObjPool, dbObj)
	}

	tx, err := urlRepo.db.Begin()
	if err != nil {
		return errors.Wrap(err, "repo.DeleteBatch")
	}
	defer tx.Rollback()

	// prepare each object to commit.
	for _, dbObj := range dbObjPool {
		_, err := tx.Model(&dbObj).
			WherePK().
			//Where("user_uuid = ?", userUUID.String()).
			Where(notDeleted).
			Update()
		if err != nil {
			return errors.Wrap(err, "repo.DeleteBatch")
		}
	}

	// commit on success.
	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "repo.DeleteBatch")
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

//EnqueueDelete uses to add batch codes to the queue chan.
func (urlRepo *URLPgRepo) EnqueueDelete(codes []string, userUUID uuid.UUID) error {

	// enqueue user codes.
	go func() {
		for _, code := range codes {
			batchUserCode := model.BatchUserCode{
				Code:     code,
				UserUUID: userUUID,
			}

			urlRepo.WorkerDeleteChan <- batchUserCode
		}
	}()

	return nil
}

// batchDeleteWorker add new url codes to the queue chan
// until batchQueueSize, than delete items in the queue.
func (urlRepo *URLPgRepo) batchDeleteWorker(batchQueueSize int) {
	go func() {
		var batchQueue []model.BatchUserCode
		ctx := context.Background()

		for itemQueue := range urlRepo.WorkerDeleteChan {
			if len(batchQueue) >= batchQueueSize {
				var entities []model.Entity
				now := time.Now()
				for _, queueItem := range batchQueue {

					// check if code created by current user.
					entity := urlRepo.checkUserCode(ctx, queueItem.Code, queueItem.UserUUID)

					// if entity exists, then remove.
					if entity != nil {
						entity.DeletedAt = &now
						entities = append(entities, *entity)
					}
				}

				if err := urlRepo.DeleteBatch(ctx, entities); err != nil {
					log.Println(err)
				}

				batchQueue = batchQueue[:0]
			}
			batchQueue = append(batchQueue, itemQueue)
		}
	}()
}

func (urlRepo *URLPgRepo) checkUserCode(ctx context.Context, code string, userUUID uuid.UUID) *model.Entity {
	entity, _ := urlRepo.GetURL(ctx, code)
	if entity != nil && entity.UserUUID == userUUID {
		return entity
	}
	return nil
}
