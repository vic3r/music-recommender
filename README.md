# Music Recommendation Engine

A Go backend for finding songs with similar vibes based on audio embeddings. Uses in-memory brute-force search with cosine similarity.

## Architecture

Clean / Hexagonal style with SOLID principles:

```
internal/
├── domain/           # Core entities (Song, SearchResult), domain errors
├── repository/       # Ports: SongRepository interface, MetadataFilter, SearchParams
├── store/            # Adapter: MemoryStore implements SongRepository
├── api/              # HTTP layer
│   ├── dto/          # Request/Response DTOs (separate from domain)
│   ├── handlers.go   # HTTP handlers (depend on repository interface)
│   ├── handler.go    # Handler struct with DI
│   ├── router.go
│   └── response.go
└── ...
```

- **DIP**: Handlers depend on `repository.SongRepository`, not concrete `store.MemoryStore`
- **SRP**: Domain, DTOs, and adapters are separate
- **ISP**: Small interfaces (SongRepository, MetadataFilter)

## Quick Start

### Run locally

```bash
go run ./cmd/server
```

### Run with Docker

```bash
docker-compose up --build
```

The API listens on `http://localhost:8080`.

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/health` | Health check |
| POST | `/api/v1/songs` | Insert a song with embedding |
| POST | `/api/v1/songs/search` | Search by embedding |
| POST | `/api/v1/songs/search/by-id` | Search similar to a song by ID |
| GET | `/api/v1/songs/{id}` | Get song metadata |
| DELETE | `/api/v1/songs/{id}` | Delete a song |

### Insert a song

```bash
curl -X POST http://localhost:8080/api/v1/songs \
  -H "Content-Type: application/json" \
  -d '{
    "embedding": [0.1, -0.2, 0.3, ...],
    "metadata": {"title": "Song Name", "artist": "Artist", "genre": "pop"}
  }'
```

### Search by embedding

```bash
curl -X POST http://localhost:8080/api/v1/songs/search \
  -H "Content-Type: application/json" \
  -d '{
    "embedding": [0.1, -0.2, 0.3, ...],
    "k": 10,
    "filter": {"genre": "pop"}
  }'
```

### Search similar to a song

```bash
curl -X POST http://localhost:8080/api/v1/songs/search/by-id \
  -H "Content-Type: application/json" \
  -d '{"id": "<song-id>", "k": 5}'
```

### Delete a song

```bash
curl -X DELETE http://localhost:8080/api/v1/songs/{id}
```

## Configuration

| Env Var | Default | Description |
|---------|---------|-------------|
| `PORT` | `8080` | HTTP port |
| `EMBEDDING_DIM` | `128` | Embedding vector dimension |

## Roadmap

- [x] **MVP**: In-memory flat search with cosine similarity
- [ ] **Indexing**: HNSW for approximate nearest neighbor
- [ ] **Persistence**: Store vectors and index on disk
- [ ] **Metadata filtering**: Combine vector search with filters (partial)
- [ ] **Scalability**: Sharding, replication, distributed search
# music-recommender
