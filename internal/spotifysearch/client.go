package spotifysearch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// TrackWithFeatures is a track with metadata and embedding from Spotify audio features.
type TrackWithFeatures struct {
	ID        string            `json:"id"`
	Embedding []float32         `json:"embedding,omitempty"`
	Metadata  map[string]string `json:"metadata,omitempty"`
}

// TracksResponse is the response from GET /api/v1/tracks/with-features (Rust SearchResponse).
type TracksResponse struct {
	Tracks []TrackWithFeatures `json:"tracks"`
	Total  int                 `json:"total"`
	Limit  int                 `json:"limit"`
	Offset int                 `json:"offset"`
}

// Client calls the Rust Spotify search service.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new Spotify search service client.
func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    strings.TrimSuffix(baseURL, "/"),
		httpClient: &http.Client{},
	}
}

// GetTracksWithFeatures fetches tracks by IDs with metadata and embeddings.
func (c *Client) GetTracksWithFeatures(trackIDs []string) (*TracksResponse, error) {
	if len(trackIDs) == 0 {
		return &TracksResponse{Tracks: []TrackWithFeatures{}}, nil
	}

	idsParam := strings.Join(trackIDs, ",")
	u := c.baseURL + "/api/v1/tracks/with-features?ids=" + url.QueryEscape(idsParam)

	resp, err := c.httpClient.Get(u)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errBody struct {
			Error string `json:"error"`
		}
		_ = json.NewDecoder(resp.Body).Decode(&errBody)
		return nil, fmt.Errorf("spotify service error %d: %s", resp.StatusCode, errBody.Error)
	}

	var tr TracksResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	return &tr, nil
}
