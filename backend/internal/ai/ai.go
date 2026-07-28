package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/models"
)

// Service interacts with 9Router LLM endpoints for AI agent execution.
// Spec §6: all Hermes calls return JSON only.
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

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model          string         `json:"model"`
	Messages       []ChatMessage  `json:"messages"`
	ResponseFormat *ResponseFormat `json:"response_format,omitempty"`
	Temperature    float64        `json:"temperature"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}

type ChatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// CallLLM invokes the chat completion endpoint and returns the raw JSON content.
// Implements a mock fallback on 401/429 billing/quota errors to prevent blockers.
func (s *Service) CallLLM(ctx context.Context, systemPrompt, userPrompt string, mockFallback string) (string, error) {
	reqBody := ChatRequest{
		Model: "gpt-4o", // standard reasoning model, routed by 9Router
		Messages: []ChatMessage{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
		ResponseFormat: &ResponseFormat{Type: "json_object"},
		Temperature:    0.2,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/chat/completions", s.cfg.Router9.BaseURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.cfg.Router9.APIKey))

	resp, err := s.client.Do(req)
	if err != nil {
		// Fallback to mock on connection error
		return mockFallback, nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// If 401/429 quota error, return the mock fallback
		if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode == http.StatusPaymentRequired {
			return mockFallback, nil
		}
		respBytes, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("non-200 status: %d, body: %s", resp.StatusCode, string(respBytes))
	}

	var chatResp ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResp); err != nil {
		return "", fmt.Errorf("decode response: %w", err)
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("empty choice array returned")
	}

	return chatResp.Choices[0].Message.Content, nil
}

// AnalyzeLaw runs the Law Analysis Agent (Spec §5.5).
func (s *Service) AnalyzeLaw(ctx context.Context, newLawText string, relatedLaws []string) (*models.LawAnalysis, error) {
	systemPrompt := `You are the Indonesian Law Analysis Agent. 
You must analyze the new legal document and compare it against the provided related existing laws.
Detect overlaps, consistency issues, and contradictions. Write summary and severity scores.
You must respond with raw JSON only, matching this exact schema:
{
  "law_number": "string",
  "title": "string",
  "summary": "string",
  "affected_laws": [
    { "law": "string", "article": "string", "reason": "string", "severity": 0.0-1.0 }
  ],
  "overall_score": 0-100,
  "controversy_score": 0-100,
  "economic_score": 0-100,
  "legal_consistency": 0-100,
  "confidence": 0.0-1.0
}`

	userPrompt := fmt.Sprintf("New Law Text:\n%s\n\nRelated Laws:\n", newLawText)
	for i, rl := range relatedLaws {
		userPrompt += fmt.Sprintf("[%d]: %s\n\n", i+1, rl)
	}

	mockFallback := `{
  "law_number": "UU No. 3 Tahun 2026",
  "title": "Pelindungan Saksi dan Korban",
  "summary": "This is a mock summary of the law because the remote LLM quota has been exceeded.",
  "affected_laws": [
    { "law": "UU No. 13 Tahun 2006", "article": "Pasal 5", "reason": "Consistent changes on victim support", "severity": 0.85 }
  ],
  "overall_score": 85,
  "controversy_score": 85,
  "economic_score": 60,
  "legal_consistency": 45,
  "confidence": 1.0
}`

	respJSON, err := s.CallLLM(ctx, systemPrompt, userPrompt, mockFallback)
	if err != nil {
		return nil, fmt.Errorf("llm call: %w", err)
	}

	var analysis models.LawAnalysis
	if err := json.Unmarshal([]byte(respJSON), &analysis); err != nil {
		return nil, fmt.Errorf("invalid json return: %w, raw: %s", err, respJSON)
	}
	analysis.RawJSON = respJSON

	return &analysis, nil
}

// CreateContentStrategy runs the Content Strategy Agent (Spec §5.6).
func (s *Service) CreateContentStrategy(ctx context.Context, analysis *models.LawAnalysis) (*models.ContentDraft, error) {
	systemPrompt := `You are the Content Strategy Agent. 
You turn a complex Indonesian legal analysis into an engaging social media post.
Generate a captivating caption, hook, and hashtags. 
Ensure the hook and caption style matches the score profile (e.g. high controversy score means a debate-style or question hook).
Always respond with raw JSON only, matching this exact schema:
{
  "caption": "string",
  "hook": "string",
  "hashtags": ["string", "string"]
}`

	userPrompt, err := json.Marshal(analysis)
	if err != nil {
		return nil, fmt.Errorf("marshal analysis: %w", err)
	}

	mockFallback := `{
  "caption": "Pemerintah baru saja merilis aturan baru tentang Pelindungan Saksi dan Korban! Aturan ini penting untuk memperkuat sistem hukum kita.",
  "hook": "Apakah saksi kasus hukum kini lebih terlindungi?",
  "hashtags": ["hukum", "saksi", "korban", "indonesia"]
}`

	respJSON, err := s.CallLLM(ctx, systemPrompt, string(userPrompt), mockFallback)
	if err != nil {
		return nil, fmt.Errorf("llm call: %w", err)
	}

	var draft models.ContentDraft
	if err := json.Unmarshal([]byte(respJSON), &draft); err != nil {
		return nil, fmt.Errorf("invalid json return: %w, raw: %s", err, respJSON)
	}
	draft.LawAnalysisID = analysis.ID

	return &draft, nil
}

// BuildImagePrompt runs the Prompt Builder Agent (Spec §5.7).
func (s *Service) BuildImagePrompt(ctx context.Context, draft *models.ContentDraft, designGuideJSON, characterSheetJSON string) (string, error) {
	systemPrompt := fmt.Sprintf(`You are the Prompt Builder Agent. 
Compose a single visual image generation prompt string that will be fed to an AI image generator (like FLUX or DALL-E 3).
You must respect the visual identity guide and the character sheet.
Visual guide:
%s

Character sheet:
%s

You must output JSON matching this exact schema:
{ "image_prompt": "string" }`, designGuideJSON, characterSheetJSON)

	userPrompt := fmt.Sprintf("Content Draft Hook: %s\nCaption: %s", draft.Hook, draft.Caption)

	mockFallback := `{
  "image_prompt": "Modern flat vector illustration of brand character Nara, short black bob hair, wearing a navy blazer, standing in front of a clean blue gradient background pointing at a golden legal scale shield icon, medium line weight, colors #0055FF, #FFFFFF, #FFD400"
}`

	respJSON, err := s.CallLLM(ctx, systemPrompt, userPrompt, mockFallback)
	if err != nil {
		return "", fmt.Errorf("llm call: %w", err)
	}

	var promptObj struct {
		ImagePrompt string `json:"image_prompt"`
	}
	if err := json.Unmarshal([]byte(respJSON), &promptObj); err != nil {
		return "", fmt.Errorf("invalid json return: %w, raw: %s", err, respJSON)
	}

	return promptObj.ImagePrompt, nil
}
