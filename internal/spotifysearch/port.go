package spotifysearch

// TracksFetcher fetches tracks with embeddings from an external service (Rust / Spotify).
type TracksFetcher interface {
	GetTracksWithFeatures(trackIDs []string) (*TracksResponse, error)
}
