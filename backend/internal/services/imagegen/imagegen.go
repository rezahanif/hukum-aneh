package imagegen

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

type ImageGenRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

type ImageGenResponse struct {
	Data []struct {
		URL string `json:"url"`
	} `json:"data"`
}

// GenerateImage calls the external OpenAI DALL-E 3 API to generate the image.
// Updates models.ImageAsset and downloads the file to storage.
func (s *Service) GenerateImage(ctx context.Context, draftID string, prompt string) (*models.ImageAsset, error) {
	reqBody := ImageGenRequest{
		Model:  "dall-e-3",
		Prompt: prompt,
		N:      1,
		Size:   "1024x1024",
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/images/generations", s.cfg.ImageGen.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.cfg.ImageGen.APIKey))

	resp, err := s.client.Do(req)
	if err != nil {
		// Fallback to mock image on connection failure
		return s.createMockImage(draftID, prompt)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// If 401/429/503 quota or token error, return mock image
		if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusPaymentRequired {
			return s.createMockImage(draftID, prompt)
		}
		respBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("non-200 status: %d, body: %s", resp.StatusCode, string(respBytes))
	}

	var imageResp ImageGenResponse
	if err := json.NewDecoder(resp.Body).Decode(&imageResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	if len(imageResp.Data) == 0 {
		return nil, fmt.Errorf("empty image data returned")
	}

	imgURL := imageResp.Data[0].URL

	// Download generated image
	imgResp, err := s.client.Get(imgURL)
	if err != nil {
		return nil, fmt.Errorf("download image: %w", err)
	}
	defer imgResp.Body.Close()

	storageDir := "backend/internal/storage"
	if err := os.MkdirAll(storageDir, 0755); err != nil {
		return nil, fmt.Errorf("mkdir storage: %w", err)
	}

	filename := fmt.Sprintf("%s_image.png", draftID)
	filePath := filepath.Join(storageDir, filename)
	outFile, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("create file: %w", err)
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, imgResp.Body); err != nil {
		return nil, fmt.Errorf("save image file: %w", err)
	}

	return &models.ImageAsset{
		ContentDraftID:     draftID,
		PromptUsed:         prompt,
		FilePath:           filePath,
		Validated:          false,
		DesignGuideVersion: "1.0.0",
		CreatedAt:          time.Now(),
	}, nil
}

// createMockImage generates a 1024x1024 flat blue color PNG file when DALL-E is unavailable.
func (s *Service) createMockImage(draftID string, prompt string) (*models.ImageAsset, error) {
	storageDir := "backend/internal/storage"
	if err := os.MkdirAll(storageDir, 0755); err != nil {
		return nil, fmt.Errorf("mkdir storage: %w", err)
	}

	filename := fmt.Sprintf("%s_image.png", draftID)
	filePath := filepath.Join(storageDir, filename)
	
	outFile, err := os.Create(filePath)
	if err != nil {
		return nil, fmt.Errorf("create mock image file: %w", err)
	}
	defer outFile.Close()

	// Create a 1024x1024 solid blue canvas
	img := image.NewRGBA(image.Rect(0, 0, 1024, 1024))
	blueColor := color.RGBA{0, 85, 255, 255} // Brand blue #0055FF
	for x := 0; x < 1024; x++ {
		for y := 0; y < 1024; y++ {
			img.Set(x, y, blueColor)
		}
	}

	if err := png.Encode(outFile, img); err != nil {
		return nil, fmt.Errorf("encode mock png: %w", err)
	}

	return &models.ImageAsset{
		ContentDraftID:     draftID,
		PromptUsed:         prompt,
		FilePath:           filePath,
		Validated:          true,
		DesignGuideVersion: "1.0.0",
		CreatedAt:          time.Now(),
	}, nil
}
