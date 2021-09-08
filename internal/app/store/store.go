package store

import (
	"encoding/gob"
	"os"

	"github.com/spinel/go-musthave-shortener/internal/app/model"
)

type Store struct {
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
	memory := make(map[string]model.Entity)
	s.decoder.Decode(&memory)

	return memory, nil
}

// SaveData save memory storage to gob file
func (s *Store) SaveData(memory map[string]model.Entity) error {
	file, _ := os.Create(s.gobFileName)
	defer file.Close()
	encoder := gob.NewEncoder(file)

	return encoder.Encode(memory)
}

func (s *Store) Close() {
	s.gobFile.Close()
}
