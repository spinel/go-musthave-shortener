package pkg

import (
	"context"
	"log"
	"time"

	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
)

func NewWorkerBatchDelete(ctx context.Context, cfg config.Config, repo repository.URLStorer, ch chan *model.Entity) {
	var batchQueue []*model.Entity
	var batchQueueSize = cfg.BatchQueueSize
	var batchItemsCounter int

	for {
		batchQueue = append(batchQueue, <-ch)
		if batchItemsCounter >= batchQueueSize {
			var entities []model.Entity
			now := time.Now()
			for _, entity := range batchQueue {
				entity.DeletedAt = &now
				entities = append(entities, *entity)
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
