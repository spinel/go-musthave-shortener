package pg

import (
	"context"

	"github.com/go-pg/pg"
	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/pkg"
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
func (entityRepo *EntityPgRepo) GetEntityBy(code string) (*model.Entity, error) {
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
	return entity, nil
}

// SaveEntity - saves given model by id.
// FIXME! code not needed, just for interface compatibility as memory
func (entityRepo *EntityPgRepo) SaveEntity(code string, entity model.Entity) error {
	_, err := entityRepo.db.Model(&entity).
		Returning("*").
		Insert()
	if err != nil {
		return err
	}
	return nil
}

func (entityRepo *EntityPgRepo) GetByUser(ctx context.Context, cfg *config.Config) []model.URLMapping {
	var urlMappingPool []model.URLMapping

	return urlMappingPool
}

func (entityRepo *EntityPgRepo) GetCode(ctx context.Context, url string) (string, error) {
	var code string
	var err error

	code, err = pkg.NewGeneratedString()
	if err != nil {
		return "", err
	}

	userUUIDString := ctx.Value(model.CookieContextName).(string)
	userUUID, _ := uuid.Parse(userUUIDString)

	entity := model.Entity{
		Code:     code,
		URL:      url,
		UserUUID: userUUID,
	}

	err = entityRepo.SaveEntity(code, entity)
	if err != nil {
		return "", err
	}
	return code, nil
}

func (entityRepo *EntityPgRepo) GetMemory() model.MemoryMap {
	return nil
}
