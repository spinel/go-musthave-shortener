package store

import (
	"encoding/gob"
	"os"
	"sync"

	"github.com/spinel/go-musthave-shortener/internal/app/model"
)

type Store struct {
	mu          sync.Mutex
	gobFileName string
	gobFile     *os.File
	decoder     *gob.Decoder
}

// NewStore is a gob storage builder
func NewStore(gobFileName string) *Store {
	gobFile, _ := os.Open(gobFileName)
	return &Store{
		gobFileName: gobFileName,
		gobFile:     gobFile,
		decoder:     gob.NewDecoder(gobFile),
	}
}

// GetData retrives data from gob file
func (s *Store) GetData() (map[string]model.Entity, error) {
	s.mu.Lock()
	memory := make(map[string]model.Entity)
	s.decoder.Decode(&memory)
	defer s.mu.Unlock()

	return memory, nil
}

// SaveData save memory storage to gob file
func (s *Store) SaveData(memory map[string]model.Entity) error {
	s.mu.Lock()
	file, _ := os.Create(s.gobFileName)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	defer s.mu.Unlock()

	return encoder.Encode(memory)
}

func (s *Store) Close() {
	s.gobFile.Close()
}
