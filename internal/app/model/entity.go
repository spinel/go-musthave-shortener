package model

// Entity is stored object.
type Entity struct {
	URL string
}

// Result is response url struct.
type Result struct {
	URL string `json:"result"`
}
