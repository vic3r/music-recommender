package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Router returns the HTTP router for the API.
func Router(h *Handler) http.Handler {
	r := chi.NewRouter()

	r.Get("/health", HandleHealth)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/songs", h.HandleInsert)
		r.Post("/songs/search", h.HandleSearch)
		r.Post("/songs/search/by-id", h.HandleSearchByID)
		r.Get("/songs/{id}", h.HandleGet)
		r.Delete("/songs/{id}", h.HandleDelete)
	})

	return r
}
