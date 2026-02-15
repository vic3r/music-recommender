package api

import "github.com/music-recommender/internal/repository"

// Handler holds dependencies for HTTP handlers.
// Depends on the SongRepository interface (DIP), not concrete implementations.
type Handler struct {
	repo repository.SongRepository
}

// NewHandler creates a new API handler with the given repository.
func NewHandler(repo repository.SongRepository) *Handler {
	return &Handler{repo: repo}
}
