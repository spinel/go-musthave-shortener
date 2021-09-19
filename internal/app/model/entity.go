package model

import "github.com/google/uuid"

// Entity is stored URL object.
// TODO: use net/url enstead of string.
// Current struct used as value of memory map.
type Entity struct {
	URL      string
	UserUUID uuid.UUID
}

// Result is response url struct for API.
type Result struct {
	URL string `json:"result"`
}

type MemoryMap map[string]Entity

type URLMapping struct {
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original_url"`
}
