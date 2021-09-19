package repository

import "github.com/spinel/go-musthave-shortener/internal/app/model"

// URLShortener defines ntity operations.
type URLShortener interface {
	GetEntityBy(id string) (*model.Entity, error)
	SaveEntity(id string, entity model.Entity) error
	GetCode(url string) (string, error)
	IncludesCode(id string) bool
	GetMemory() model.MemoryMap
}
