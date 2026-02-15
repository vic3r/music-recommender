package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/music-recommender/internal/api"
	"github.com/music-recommender/internal/store"
	"github.com/music-recommender/internal/spotifysearch"
)

func main() {
	// Spotify audio features = 12 dimensions; use EMBEDDING_DIM=12 for import saga
	dim := 12
	if d := os.Getenv("EMBEDDING_DIM"); d != "" {
		if n, err := strconv.Atoi(d); err == nil && n > 0 {
			dim = n
		}
	}

	addr := ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}

	repo := store.NewMemoryStore(dim)

	var spotifyFetcher spotifysearch.TracksFetcher
	if grpcTarget := os.Getenv("SPOTIFY_GRPC_TARGET"); grpcTarget != "" {
		// Prefer gRPC over HTTP
		cli, err := spotifysearch.NewGrpcClient(grpcTarget)
		if err != nil {
			log.Fatalf("spotify gRPC client: %v", err)
		}
		spotifyFetcher = cli
	} else if url := os.Getenv("SPOTIFY_SEARCH_URL"); url != "" {
		spotifyFetcher = spotifysearch.NewClient(url)
	}

	h := api.NewHandler(repo, spotifyFetcher)
	r := api.Router(h)

	log.Printf("starting music recommender server on %s (embedding_dim=%d)", addr, dim)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
