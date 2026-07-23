package publishing

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/models"
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

// PublishToInstagram posts the image and caption to Instagram via Meta Graph API.
// Note: Meta Graph API requires public URL for the image. If asset.FilePath is local,
// it must be uploaded to a public staging URL first (e.g. S3/Imgur).
func (s *Service) PublishToInstagram(ctx context.Context, draft *models.ContentDraft, asset *models.ImageAsset, publicImageURL string) (string, error) {
	if s.cfg.Instagram.AccessToken == "" || s.cfg.Instagram.AccountID == "" {
		return "", fmt.Errorf("instagram credentials missing")
	}

	caption := fmt.Sprintf("%s\n\n%s", draft.Hook, draft.Caption)

	// Step 1: Create media container
	containerURL := fmt.Sprintf(
		"https://graph.facebook.com/v17.0/%s/media?image_url=%s&caption=%s&access_token=%s",
		s.cfg.Instagram.AccountID,
		url.QueryEscape(publicImageURL),
		url.QueryEscape(caption),
		s.cfg.Instagram.AccessToken,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, containerURL, nil)
	if err != nil {
		return "", fmt.Errorf("new request: %w", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("create media container: %w", err)
	}
	defer resp.Body.Close()

	var createResp struct {
		ID    string `json:"id"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&createResp); err != nil {
		return "", fmt.Errorf("decode create response: %w", err)
	}
	if createResp.Error != nil {
		return "", fmt.Errorf("facebook api error: %s", createResp.Error.Message)
	}

	containerID := createResp.ID

	// Step 2: Poll container status
	statusURL := fmt.Sprintf(
		"https://graph.facebook.com/v17.0/%s?fields=status_code&access_token=%s",
		containerID,
		s.cfg.Instagram.AccessToken,
	)

	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(5 * time.Second):
		}

		sReq, err := http.NewRequestWithContext(ctx, http.MethodGet, statusURL, nil)
		if err != nil {
			return "", err
		}
		sResp, err := s.client.Do(sReq)
		if err != nil {
			continue
		}

		var status struct {
			StatusCode string `json:"status_code"`
		}
		json.NewDecoder(sResp.Body).Decode(&status)
		sResp.Body.Close()

		if status.StatusCode == "FINISHED" {
			break
		}
		if status.StatusCode == "ERROR" {
			return "", fmt.Errorf("media container status returned error")
		}
	}

	// Step 3: Publish media container
	publishURL := fmt.Sprintf(
		"https://graph.facebook.com/v17.0/%s/media_publish?creation_id=%s&access_token=%s",
		s.cfg.Instagram.AccountID,
		containerID,
		s.cfg.Instagram.AccessToken,
	)

	pReq, err := http.NewRequestWithContext(ctx, http.MethodPost, publishURL, nil)
	if err != nil {
		return "", err
	}
	pResp, err := s.client.Do(pReq)
	if err != nil {
		return "", fmt.Errorf("publish container: %w", err)
	}
	defer pResp.Body.Close()

	var publishResp struct {
		ID    string `json:"id"`
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err := json.NewDecoder(pResp.Body).Decode(&publishResp); err != nil {
		return "", fmt.Errorf("decode publish response: %w", err)
	}
	if publishResp.Error != nil {
		return "", fmt.Errorf("publish api error: %s", publishResp.Error.Message)
	}

	return publishResp.ID, nil
}
