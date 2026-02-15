package api

import (
	"encoding/json"
	"net/http"

	"github.com/music-recommender/internal/api/dto"
	"github.com/music-recommender/internal/domain"
	"github.com/music-recommender/internal/repository"
)

// HandleInsert adds a song with embedding and metadata.
func (h *Handler) HandleInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var req dto.InsertRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	song, err := h.repo.Insert(req.Embedding, req.Metadata)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, dto.InsertResponse{ID: song.ID})
}

// HandleSearch finds songs similar to the given embedding.
func (h *Handler) HandleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var req dto.SearchRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	params := buildSearchParams(req.Embedding, req.K, req.Filter)
	results, err := h.repo.Search(params)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, dto.SearchResponse{
		Results: toSearchResultItems(results),
	})
}

// HandleSearchByID finds songs similar to an existing song.
func (h *Handler) HandleSearchByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	var req dto.SearchByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondError(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	song, ok := h.repo.Get(req.ID)
	if !ok {
		respondError(w, http.StatusNotFound, "song not found")
		return
	}
	k := req.K
	if k <= 0 {
		k = 10
	}
	params := buildSearchParams(song.Embedding, k+1, req.Filter)
	results, err := h.repo.Search(params)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	items := excludeAndLimit(results, req.ID, k)
	respondJSON(w, http.StatusOK, dto.SearchResponse{Results: items})
}

// HandleDelete removes a song by ID.
func (h *Handler) HandleDelete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	id := r.PathValue("id")
	if id == "" {
		respondError(w, http.StatusBadRequest, "missing id")
		return
	}
	if h.repo.Delete(id) {
		w.WriteHeader(http.StatusNoContent)
	} else {
		respondError(w, http.StatusNotFound, "song not found")
	}
}

// HandleGet returns a song by ID (metadata only).
func (h *Handler) HandleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	id := r.PathValue("id")
	if id == "" {
		respondError(w, http.StatusBadRequest, "missing id")
		return
	}
	song, ok := h.repo.Get(id)
	if !ok {
		respondError(w, http.StatusNotFound, "song not found")
		return
	}
	respondJSON(w, http.StatusOK, dto.SongResponse{
		ID:       song.ID,
		Metadata: song.Metadata,
	})
}

// HandleHealth returns service health.
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]any{"status": "ok"})
}

// buildSearchParams converts DTO fields to repository.SearchParams.
func buildSearchParams(query []float32, k int, filter map[string]string) repository.SearchParams {
	params := repository.SearchParams{Query: query, K: k}
	if len(filter) > 0 {
		params.Filter = repository.KeyValueFilter(filter)
	}
	return params
}

// toSearchResultItems maps domain.SearchResult to DTOs.
func toSearchResultItems(results []domain.SearchResult) []dto.SearchResultItem {
	items := make([]dto.SearchResultItem, len(results))
	for i, r := range results {
		items[i] = dto.SearchResultItem{
			ID:       r.Song.ID,
			Metadata: r.Song.Metadata,
			Score:    r.Score,
		}
	}
	return items
}

// excludeAndLimit filters out the query song and limits to k results.
func excludeAndLimit(results []domain.SearchResult, excludeID string, k int) []dto.SearchResultItem {
	var items []dto.SearchResultItem
	for _, r := range results {
		if r.Song.ID == excludeID {
			continue
		}
		items = append(items, dto.SearchResultItem{
			ID:       r.Song.ID,
			Metadata: r.Song.Metadata,
			Score:    r.Score,
		})
		if len(items) >= k {
			break
		}
	}
	return items
}