package analytics

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
)

type Service struct {
	cfg    *config.Config
	client *http.Client
}

func New(cfg *config.Config) *Service {
	return &Service{
		cfg:    cfg,
		client: &http.Client{},
	}
}

type IGMetrics struct {
	LikeCount    int `json:"like_count"`
	CommentCount int `json:"comments_count"`
}

// CollectInstagramMetrics retrieves engagement stats for a published Instagram post.
func (s *Service) CollectInstagramMetrics(ctx context.Context, externalPostID string) (*IGMetrics, error) {
	if s.cfg.Instagram.AccessToken == "" {
		return nil, fmt.Errorf("instagram access token missing")
	}

	url := fmt.Sprintf(
		"https://graph.facebook.com/v17.0/%s?fields=like_count,comments_count&access_token=%s",
		externalPostID,
		s.cfg.Instagram.AccessToken,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("facebook api status: %d", resp.StatusCode)
	}

	var metrics IGMetrics
	if err := json.NewDecoder(resp.Body).Decode(&metrics); err != nil {
		return nil, fmt.Errorf("decode metrics: %w", err)
	}

	return &metrics, nil
}
