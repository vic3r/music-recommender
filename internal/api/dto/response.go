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

// ImportResponse is the response from the import saga.
type ImportResponse struct {
	Imported []string            `json:"imported"` // IDs of successfully imported songs (Spotify IDs)
	Failed   []string            `json:"failed,omitempty"` // IDs that failed to import
	Similar  []SearchResultItem  `json:"similar,omitempty"` // Similar tracks if find_similar_to was set
}

// ErrorResponse is the JSON error payload.
type ErrorResponse struct {
	Error string `json:"error"`
}
