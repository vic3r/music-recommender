package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/music-recommender/internal/api"
	"github.com/music-recommender/internal/store"
)

func main() {
	dim := 128 // default embedding dimension (e.g. typical for audio embeddings)
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
	h := api.NewHandler(repo)
	r := api.Router(h)

	log.Printf("starting music recommender server on %s (embedding_dim=%d)", addr, dim)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
