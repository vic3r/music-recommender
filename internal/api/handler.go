package api

import (
	"github.com/music-recommender/internal/repository"
	"github.com/music-recommender/internal/spotifysearch"
)

// Handler holds dependencies for HTTP handlers.
// Depends on interfaces (DIP), not concrete implementations.
type Handler struct {
	repo            repository.SongRepository
	spotifyFetcher  spotifysearch.TracksFetcher // nil if SPOTIFY_SEARCH_URL not set
}

// NewHandler creates a new API handler with the given repository.
func NewHandler(repo repository.SongRepository, spotifyFetcher spotifysearch.TracksFetcher) *Handler {
	return &Handler{repo: repo, spotifyFetcher: spotifyFetcher}
}
