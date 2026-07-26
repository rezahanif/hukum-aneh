package drive

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type Service struct {
	driveSrv *drive.Service
	folderID string
}

type savingTokenSource struct {
	oauth2.TokenSource
	tokenPath string
}

func (s *savingTokenSource) Token() (*oauth2.Token, error) {
	t, err := s.TokenSource.Token()
	if err != nil {
		return nil, err
	}
	_ = saveToken(s.tokenPath, t)
	return t, nil
}

func saveToken(path string, token *oauth2.Token) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(token)
}

func New(ctx context.Context, credsPath, tokenPath, folderID string) (*Service, error) {
	credBytes, err := os.ReadFile(credsPath)
	if err != nil {
		return nil, fmt.Errorf("read credentials: %w", err)
	}

	config, err := google.ConfigFromJSON(credBytes, drive.DriveFileScope)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	tokenFile, err := os.Open(tokenPath)
	if err != nil {
		return nil, fmt.Errorf("open token: %w", err)
	}
	defer tokenFile.Close()

	tok := &oauth2.Token{}
	if err := json.NewDecoder(tokenFile).Decode(tok); err != nil {
		return nil, fmt.Errorf("decode token: %w", err)
	}

	ts := &savingTokenSource{
		TokenSource: config.TokenSource(ctx, tok),
		tokenPath:   tokenPath,
	}

	client := oauth2.NewClient(ctx, ts)
	driveSrv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("create drive service: %w", err)
	}

	return &Service{
		driveSrv: driveSrv,
		folderID: folderID,
	}, nil
}

func (s *Service) UploadPDF(ctx context.Context, name string, content io.Reader) (string, error) {
	fileMetadata := &drive.File{
		Name:     name,
		MimeType: "application/pdf",
	}
	if s.folderID != "" {
		fileMetadata.Parents = []string{s.folderID}
	}

	res, err := s.driveSrv.Files.Create(fileMetadata).
		Media(content).
		Context(ctx).
		Fields("id").
		Do()
	if err != nil {
		return "", fmt.Errorf("upload to drive: %w", err)
	}

	return res.Id, nil
}
