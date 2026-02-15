package dto

// InsertResponse is the HTTP response after a successful insert.
type InsertResponse struct {
	ID string `json:"id"`
}

// SearchResponse is the HTTP response for search endpoints.
type SearchResponse struct {
	Results []SearchResultItem `json:"results"`
}

// SearchResultItem is a single search result (metadata only, no embedding).
type SearchResultItem struct {
	ID       string            `json:"id"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Score    float32           `json:"score"`
}

// SongResponse is the HTTP response for a single song (metadata only).
type SongResponse struct {
	ID       string            `json:"id"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

// ErrorResponse is the JSON error payload.
type ErrorResponse struct {
	Error string `json:"error"`
}
