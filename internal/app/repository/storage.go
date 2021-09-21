package repository

import (
	"encoding/gob"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/spinel/go-musthave-shortener/internal/app/config"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/repository/pg"
	"github.com/spinel/go-musthave-shortener/internal/app/repository/web"
)

type Storage struct {
	mu          sync.Mutex
	gobFileName string
	gobFile     *os.File

	Pg       *pg.DB
	EntityPg URLShortener

	Entity URLShortener
}

// NewStorage is a gob storage builder
func NewStorage(cfg *config.Config) (*Storage, error) {
	pgDB, err := pg.Dial(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "pgdb.Dial failed")
	}

	gobFile, _ := os.Open(cfg.GobFileName)
	memory := make(model.MemoryMap)
	decoder := gob.NewDecoder(gobFile)
	decoder.Decode(&memory)

	entityRepo := web.NewEntityRepo(memory)
	entityRepoPg := pg.NewEntityPgRepo(pgDB)

	s := &Storage{
		gobFileName: cfg.GobFileName,
		gobFile:     gobFile,

		Entity: entityRepo,

		Pg:       pgDB,
		EntityPg: entityRepoPg,
	}
	s.flush()

	return s, nil
}

// SaveData save memory storage to gob file
func (s *Storage) saveData(memory model.MemoryMap) error {
	s.mu.Lock()
	file, _ := os.Create(s.gobFileName)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	defer s.mu.Unlock()

	return encoder.Encode(memory)
}

func (s *Storage) close() {
	s.gobFile.Close()
}

func (s *Storage) flush() {
	ticker := time.NewTicker(5 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				memory := s.Entity.GetMemory()
				err := s.saveData(memory)
				if err != nil {
					panic(err)
				}
				fmt.Println("memory flushed at", t)
			}
		}
	}()

}
