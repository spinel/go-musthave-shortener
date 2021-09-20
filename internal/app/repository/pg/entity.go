package pg

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
