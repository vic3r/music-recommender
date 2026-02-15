package spotifysearch

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	spotifypb "github.com/music-recommender/internal/proto"
)

// GrpcClient calls the Rust Spotify search service via gRPC.
type GrpcClient struct {
	conn   *grpc.ClientConn
	client spotifypb.SpotifySearchClient
}

// NewGrpcClient creates a gRPC client for the Spotify search service.
func NewGrpcClient(target string) (*GrpcClient, error) {
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc dial: %w", err)
	}
	return &GrpcClient{
		conn:   conn,
		client: spotifypb.NewSpotifySearchClient(conn),
	}, nil
}

// Close closes the gRPC connection.
func (c *GrpcClient) Close() error {
	return c.conn.Close()
}

// GetTracksWithFeatures fetches tracks by IDs with metadata and embeddings.
func (c *GrpcClient) GetTracksWithFeatures(trackIDs []string) (*TracksResponse, error) {
	if len(trackIDs) == 0 {
		return &TracksResponse{Tracks: []TrackWithFeatures{}}, nil
	}

	resp, err := c.client.GetTracksWithFeatures(context.Background(), &spotifypb.GetTracksWithFeaturesRequest{
		TrackIds: trackIDs,
	})
	if err != nil {
		return nil, fmt.Errorf("grpc GetTracksWithFeatures: %w", err)
	}

	tracks := make([]TrackWithFeatures, len(resp.Tracks))
	for i, t := range resp.Tracks {
		tracks[i] = TrackWithFeatures{
			ID:        t.Id,
			Embedding: t.Embedding,
			Metadata:  t.Metadata,
		}
	}

	return &TracksResponse{
		Tracks: tracks,
		Total:  len(tracks),
	}, nil
}

// Ensure GrpcClient implements TracksFetcher
var _ TracksFetcher = (*GrpcClient)(nil)
