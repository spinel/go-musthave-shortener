package model

// Entity is stored URL object.
// TODO: use net/url enstead of string.
// Current struct used as value of memory map.
type Entity struct {
	URL string
}

// Result is response url struct for API.
type Result struct {
	URL string `json:"result"`
}

type MemoryMap map[string]Entity
