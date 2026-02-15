package domain

// Song is the core domain entity for a track with its embedding and metadata.
// Embedding is excluded from JSON serialization for API responses.
type Song struct {
	ID        string
	Embedding []float32
	Metadata  map[string]string
}
