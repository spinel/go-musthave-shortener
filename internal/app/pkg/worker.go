package pkg

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
)

func NewWorkerBatchDelete(ctx context.Context, cfg config.Config, repo repository.URLStorer, ch chan *model.BatchUserCode) {
	var batchQueue []*model.BatchUserCode
	var batchQueueSize = cfg.BatchQueueSize
	var batchItemsCounter int

	for {
		batchQueue = append(batchQueue, <-ch)
		if batchItemsCounter >= batchQueueSize {
			var entities []model.Entity
			now := time.Now()
			for _, queueItem := range batchQueue {

				// check if code created by current user.
				entity := checkUserCode(ctx, repo, queueItem.Code, queueItem.UserUUID)

				// if entity exists, then remove.
				if entity != nil {
					entity.DeletedAt = &now
					entities = append(entities, *entity)
				}
			}

			err := repo.DeleteBatch(ctx, entities)
			if err != nil {
				log.Println(err)
			}

			batchQueue = batchQueue[:0]
			batchItemsCounter = 0
		}
		batchItemsCounter++
	}
}

func checkUserCode(ctx context.Context, repo repository.URLStorer, code string, userUUID uuid.UUID) *model.Entity {
	entity, _ := repo.GetURL(ctx, code)
	if entity != nil && entity.UserUUID == userUUID {
		return entity
	}
	return nil
}
