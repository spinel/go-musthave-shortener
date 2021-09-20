package pg

import (
	"errors"

	"github.com/go-pg/pg"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
)

type EntityPgRepo struct {
	db *DB
}

func NewEntityPgRepo(db *DB) *EntityPgRepo {
	return &EntityPgRepo{db: db}
}

func (entityRepo *EntityPgRepo) Ping() bool {
	_, err := entityRepo.db.Exec("SELECT 1")
	return err == nil
}

// GetEntityBy - retrive entity by id.
func (entityRepo *EntityPgRepo) GetEntityBy(id string) (*model.Entity, error) {
	entity := &model.Entity{}
	err := entityRepo.db.Model(entity).
		Where("code = ?", code).
		Where(notDeleted).
		Select()
	if err != nil {
		if err.Error() == pg.ErrNoRows.Error() { //not found
			return nil, nil
		}
		return nil, err
	}
	return item, nil
	return nil, errors.New("not found")
}
