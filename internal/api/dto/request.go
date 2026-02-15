package dto

// InsertRequest is the HTTP request body for adding a song.
type InsertRequest struct {
	Embedding []float32         `json:"embedding"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// SearchRequest is the HTTP request body for similarity search by embedding.
type SearchRequest struct {
	Embedding []float32         `json:"embedding"`
	K         int               `json:"k,omitempty"`
	Filter    map[string]string `json:"filter,omitempty"`
}

// SearchByIDRequest is the HTTP request body for finding songs similar to an existing song.
type SearchByIDRequest struct {
	ID     string            `json:"id"`
	K      int               `json:"k,omitempty"`
	Filter map[string]string `json:"filter,omitempty"`
}
