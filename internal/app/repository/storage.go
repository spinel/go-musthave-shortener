package repository

import (
	"encoding/gob"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/repository/web"
)

type Storage struct {
	mu          sync.Mutex
	gobFileName string
	gobFile     *os.File

	Entity URLShortener
}

// NewStorage is a gob storage builder
func NewStorage(gobFileName string) *Storage {
	gobFile, _ := os.Open(gobFileName)
	memory := make(model.MemoryMap)
	decoder := gob.NewDecoder(gobFile)
	decoder.Decode(&memory)

	// Entity interface
	entityRepo := web.NewEntityRepo(memory)

	s := &Storage{
		gobFileName: gobFileName,
		gobFile:     gobFile,

		Entity: entityRepo,
	}
	s.flush()

	return s
}

// SaveData save memory storage to gob file
func (s *Storage) SaveData(memory model.MemoryMap) error {
	s.mu.Lock()
	file, _ := os.Create(s.gobFileName)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	defer s.mu.Unlock()

	return encoder.Encode(memory)
}

func (s *Storage) Close() {
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
				err := s.SaveData(memory)
				if err != nil {
					panic(err)
				}
				fmt.Println("memory flushed at", t)
			}
		}
	}()

}
