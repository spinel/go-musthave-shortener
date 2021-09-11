package store

import (
	"encoding/gob"
	"os"
	"sync"

	"github.com/spinel/go-musthave-shortener/internal/app/model"
)

type store struct {
	mu          sync.Mutex
	gobFileName string
	gobFile     *os.File
	decoder     *gob.Decoder
}

// NewStore is a gob storage builder
func NewStore(gobFileName string) *store {
	gobFile, _ := os.Open(gobFileName)
	return &store{
		gobFileName: gobFileName,
		gobFile:     gobFile,
		decoder:     gob.NewDecoder(gobFile),
	}
}

// GetData retrives data from gob file
func (s *store) GetData() (model.MemoryMap, error) {
	s.mu.Lock()
	memory := make(model.MemoryMap)
	s.decoder.Decode(&memory)
	defer s.mu.Unlock()

	return memory, nil
}

// SaveData save memory storage to gob file
func (s *store) SaveData(memory model.MemoryMap) error {
	s.mu.Lock()
	file, _ := os.Create(s.gobFileName)
	defer file.Close()
	encoder := gob.NewEncoder(file)
	defer s.mu.Unlock()

	return encoder.Encode(memory)
}

func (s *store) Close() {
	s.gobFile.Close()
}
