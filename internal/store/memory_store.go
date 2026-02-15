package store

import (
	"sort"
	"sync"

	"github.com/google/uuid"
	"github.com/music-recommender/internal/domain"
	"github.com/music-recommender/internal/repository"
)

// MemoryStore is an in-memory implementation of repository.SongRepository.
// Uses brute-force cosine similarity search.
type MemoryStore struct {
	mu    sync.RWMutex
	songs map[string]*domain.Song
	dim   int
}

// NewMemoryStore creates a new in-memory store with the given embedding dimension.
func NewMemoryStore(dim int) *MemoryStore {
	return &MemoryStore{
		songs: make(map[string]*domain.Song),
		dim:   dim,
	}
}

// Insert adds a song with its embedding and metadata.
func (s *MemoryStore) Insert(embedding []float32, metadata map[string]string) (*domain.Song, error) {
	if len(embedding) != s.dim {
		return nil, domain.ErrDimensionMismatch
	}
	return s.insert(uuid.New().String(), embedding, metadata)
}

// InsertWithID adds a song with a specific ID (e.g. for re-insertion).
func (s *MemoryStore) InsertWithID(id string, embedding []float32, metadata map[string]string) (*domain.Song, error) {
	if len(embedding) != s.dim {
		return nil, domain.ErrDimensionMismatch
	}
	return s.insert(id, embedding, metadata)
}

func (s *MemoryStore) insert(id string, embedding []float32, metadata map[string]string) (*domain.Song, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	meta := metadata
	if meta == nil {
		meta = make(map[string]string)
	}
	song := &domain.Song{
		ID:        id,
		Embedding: embedding,
		Metadata:  meta,
	}
	s.songs[id] = song
	return song, nil
}

// Search finds the k nearest neighbors by cosine similarity (brute force).
func (s *MemoryStore) Search(params repository.SearchParams) ([]domain.SearchResult, error) {
	if len(params.Query) != s.dim {
		return nil, domain.ErrDimensionMismatch
	}

	s.mu.RLock()
	defer s.mu.RUnlock()

	k := params.K
	if k <= 0 {
		k = 10
	}

	var results []domain.SearchResult
	for _, song := range s.songs {
		if params.Filter != nil && !params.Filter.Match(song.Metadata) {
			continue
		}
		score := CosineSimilarity(params.Query, song.Embedding)
		results = append(results, domain.SearchResult{Song: song, Score: score})
	}

	sort.Slice(results, func(i, j int) bool { return results[i].Score > results[j].Score })
	if len(results) > k {
		results = results[:k]
	}
	return results, nil
}

// Get returns a song by ID.
func (s *MemoryStore) Get(id string) (*domain.Song, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	song, ok := s.songs[id]
	return song, ok
}

// Delete removes a song by ID.
func (s *MemoryStore) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.songs[id]; ok {
		delete(s.songs, id)
		return true
	}
	return false
}

// Count returns the number of songs in the store.
func (s *MemoryStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.songs)
}
