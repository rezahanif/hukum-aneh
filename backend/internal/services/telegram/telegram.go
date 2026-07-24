package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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

// InlineKeyboardButton is a Telegram inline keyboard button.
type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

// InlineKeyboardMarkup is a Telegram inline keyboard markup.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// SendApprovalRequest sends a photo with analysis info and inline buttons for approval.
func (s *Service) SendApprovalRequest(
	ctx context.Context,
	draft *models.ContentDraft,
	asset *models.ImageAsset,
	analysis *models.LawAnalysis,
	title string,
) (int, error) {
	filePath := asset.FilePath
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("open photo: %w", err)
	}
	defer file.Close()

	captionText := fmt.Sprintf(
		"⚖️ *NEW LAW DETECTED*\n\n"+
			"*Law:* %s\n"+
			"*Title:* %s\n\n"+
			"*Hook:* %s\n\n"+
			"*Caption:*\n%s\n\n"+
			"*Scores:*\n"+
			"• Overall: %d\n"+
			"• Controversy: %d\n"+
			"• Economic: %d\n"+
			"• Consistency: %d\n\n"+
			"*Hashtags:* %s",
		analysis.LawDocumentID,
		title,
		draft.Hook,
		draft.Caption,
		analysis.OverallScore,
		analysis.ControversyScore,
		analysis.EconomicScore,
		analysis.LegalConsistency,
		draft.Hashtags,
	)

	// Ensure caption is within Telegram's 1024 char limit for photo captions
	if len(captionText) > 1024 {
		captionText = captionText[:1020] + "..."
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("photo", filepath.Base(filePath))
	if err != nil {
		return 0, fmt.Errorf("create form file: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return 0, fmt.Errorf("copy file: %w", err)
	}

	_ = writer.WriteField("chat_id", s.cfg.Telegram.ChatID)
	_ = writer.WriteField("caption", captionText)
	_ = writer.WriteField("parse_mode", "Markdown")

	// Inline keyboard markup
	keyboard := InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{
			{
				{Text: "✅ Approve", CallbackData: fmt.Sprintf("approve:%s", draft.ID)},
				{Text: "❌ Reject", CallbackData: fmt.Sprintf("reject:%s", draft.ID)},
			},
			{
				{Text: "🔄 Regene Image", CallbackData: fmt.Sprintf("regen_img:%s", draft.ID)},
				{Text: "📝 Regene Caption", CallbackData: fmt.Sprintf("regen_cap:%s", draft.ID)},
			},
		},
	}
	kbdBytes, err := json.Marshal(keyboard)
	if err == nil {
		_ = writer.WriteField("reply_markup", string(kbdBytes))
	}

	if err := writer.Close(); err != nil {
		return 0, fmt.Errorf("close writer: %w", err)
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendPhoto", s.cfg.Telegram.BotToken)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return 0, fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBytes, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("non-200 status: %d, body: %s", resp.StatusCode, string(respBytes))
	}

	var tgResp struct {
		Ok     bool `json:"ok"`
		Result struct {
			MessageID int `json:"message_id"`
		} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tgResp); err != nil {
		return 0, fmt.Errorf("decode response: %w", err)
	}

	return tgResp.Result.MessageID, nil
}

// SendPromptApproval sends a text message with the law details and planned image prompt,
// with approve/reject inline buttons. This is the pre-generation approval gate.
func (s *Service) SendPromptApproval(
	ctx context.Context,
	draft *models.ContentDraft,
	analysis *models.LawAnalysis,
	title string,
) (int, error) {
	text := fmt.Sprintf(
		"🎨 *PROMPT APPROVAL NEEDED*\n\n"+
			"*Law:* %s\n"+
			"*Title:* %s\n\n"+
			"*Hook:* %s\n\n"+
			"*Caption:*\n%s\n\n"+
			"*Planned Image Prompt:*\n%s\n\n"+
			"*Scores:* Overall=%d | Controversy=%d | Economic=%d | Consistency=%d\n\n"+
			"*Hashtags:* %s",
		analysis.LawDocumentID,
		title,
		draft.Hook,
		draft.Caption,
		draft.ImagePrompt,
		analysis.OverallScore,
		analysis.ControversyScore,
		analysis.EconomicScore,
		analysis.LegalConsistency,
		draft.Hashtags,
	)

	if len(text) > 4096 {
		text = text[:4090] + "..."
	}

	keyboard := InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{
			{
				{Text: "✅ Approve Prompt", CallbackData: fmt.Sprintf("prompt_approve:%s", draft.ID)},
				{Text: "❌ Reject Prompt", CallbackData: fmt.Sprintf("prompt_reject:%s", draft.ID)},
			},
		},
	}
	kbdBytes, _ := json.Marshal(keyboard)

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", s.cfg.Telegram.BotToken)
	payload := map[string]string{
		"chat_id":      s.cfg.Telegram.ChatID,
		"text":         text,
		"parse_mode":   "Markdown",
		"reply_markup": string(kbdBytes),
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return 0, fmt.Errorf("marshal payload: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return 0, fmt.Errorf("new request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBytes, _ := io.ReadAll(resp.Body)
		return 0, fmt.Errorf("non-200 status: %d, body: %s", resp.StatusCode, string(respBytes))
	}

	var tgResp struct {
		Ok     bool `json:"ok"`
		Result struct {
			MessageID int `json:"message_id"`
		} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tgResp); err != nil {
		return 0, fmt.Errorf("decode response: %w", err)
	}

	return tgResp.Result.MessageID, nil
}

type TGUpdate struct {
	UpdateID      int `json:"update_id"`
	CallbackQuery *struct {
		ID   string `json:"id"`
		Data string `json:"data"`
		From struct {
			ID int `json:"id"`
		} `json:"from"`
	} `json:"callback_query"`
}

// StartPolling listens for callback query updates from the Telegram bot.
// Uses long polling on getUpdates. Calls onAction when callback query is received.
func (s *Service) StartPolling(ctx context.Context, onAction func(draftID string, action string, reviewerID string) error) error {
	offset := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?timeout=30&offset=%d", s.cfg.Telegram.BotToken, offset)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		resp, err := s.client.Do(req)
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		var updateResp struct {
			Ok     bool       `json:"ok"`
			Result []TGUpdate `json:"result"`
		}
		err = json.NewDecoder(resp.Body).Decode(&updateResp)
		resp.Body.Close()
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		if !updateResp.Ok {
			time.Sleep(5 * time.Second)
			continue
		}

		for _, up := range updateResp.Result {
			offset = up.UpdateID + 1

			if up.CallbackQuery != nil {
				data := up.CallbackQuery.Data
				// Format: action:draftID
				var action, draftID string
				// Data format might have colon splitter
				parts := strings.Split(data, ":")
				if len(parts) == 2 {
					action = parts[0]
					draftID = parts[1]
				}

				if draftID != "" && action != "" {
					reviewerID := fmt.Sprintf("%d", up.CallbackQuery.From.ID)
					if err := onAction(draftID, action, reviewerID); err != nil {
						s.answerCallback(up.CallbackQuery.ID, fmt.Sprintf("Error: %s", err))
					} else {
						s.answerCallback(up.CallbackQuery.ID, fmt.Sprintf("Processed: %s", action))
					}
				}
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func (s *Service) answerCallback(callbackQueryID string, text string) {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/answerCallbackQuery", s.cfg.Telegram.BotToken)
	reqBody := map[string]string{
		"callback_query_id": callbackQueryID,
		"text":              text,
	}
	bodyBytes, _ := json.Marshal(reqBody)
	_, _ = s.client.Post(url, "application/json", bytes.NewReader(bodyBytes))
}
