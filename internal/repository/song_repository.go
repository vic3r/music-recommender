package repository

import "github.com/music-recommender/internal/domain"

// MetadataFilter filters songs by metadata key-value pairs.
// Implementations decide matching logic (e.g., exact, prefix).
type MetadataFilter interface {
	Match(metadata map[string]string) bool
}

// SearchParams holds parameters for vector similarity search.
type SearchParams struct {
	Query  []float32
	K      int
	Filter MetadataFilter
}

// SongRepository defines the port for song storage and retrieval.
// Implementations can be in-memory, disk-backed, or distributed.
type SongRepository interface {
	Insert(embedding []float32, metadata map[string]string) (*domain.Song, error)
	InsertWithID(id string, embedding []float32, metadata map[string]string) (*domain.Song, error)
	Search(params SearchParams) ([]domain.SearchResult, error)
	Get(id string) (*domain.Song, bool)
	Delete(id string) bool
	Count() int
}
