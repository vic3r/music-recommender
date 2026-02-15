package api

import (
	"encoding/json"
	"net/http"

	"github.com/music-recommender/internal/api/dto"
)

func respondJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if body != nil {
		_ = json.NewEncoder(w).Encode(body)
	}
}

func respondError(w http.ResponseWriter, status int, msg string) {
	respondJSON(w, status, dto.ErrorResponse{Error: msg})
}
