package domain

// SearchResult pairs a song with its similarity score from vector search.
type SearchResult struct {
	Song  *Song
	Score float32
}
