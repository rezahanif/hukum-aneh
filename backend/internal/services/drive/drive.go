package drive

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type Service struct {
	driveSrv *drive.Service
	folderID string
	folders  map[string]string
	mu       sync.Mutex
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
		folders:  make(map[string]string),
	}, nil
}

func (s *Service) GetOrCreateFolder(ctx context.Context, name string) (string, error) {
	if name == "" {
		return s.folderID, nil
	}

	s.mu.Lock()
	if id, ok := s.folders[name]; ok {
		s.mu.Unlock()
		return id, nil
	}
	s.mu.Unlock()

	// Search for existing folder
	query := fmt.Sprintf("name = '%s' and mimeType = 'application/vnd.google-apps.folder' and trashed = false", name)
	if s.folderID != "" {
		query += fmt.Sprintf(" and '%s' in parents", s.folderID)
	}

	list, err := s.driveSrv.Files.List().
		Q(query).
		Context(ctx).
		Fields("files(id)").
		Do()
	if err != nil {
		return "", fmt.Errorf("search folder: %w", err)
	}

	if len(list.Files) > 0 {
		id := list.Files[0].Id
		s.mu.Lock()
		s.folders[name] = id
		s.mu.Unlock()
		return id, nil
	}

	// Create new folder
	folderMetadata := &drive.File{
		Name:     name,
		MimeType: "application/vnd.google-apps.folder",
	}
	if s.folderID != "" {
		folderMetadata.Parents = []string{s.folderID}
	}

	res, err := s.driveSrv.Files.Create(folderMetadata).
		Context(ctx).
		Fields("id").
		Do()
	if err != nil {
		return "", fmt.Errorf("create folder: %w", err)
	}

	s.mu.Lock()
	s.folders[name] = res.Id
	s.mu.Unlock()

	return res.Id, nil
}

func (s *Service) UploadPDF(ctx context.Context, name string, content io.Reader, group string) (string, error) {
	targetFolderID, err := s.GetOrCreateFolder(ctx, group)
	if err != nil {
		return "", fmt.Errorf("resolve group folder: %w", err)
	}

	fileMetadata := &drive.File{
		Name:     name,
		MimeType: "application/pdf",
	}
	if targetFolderID != "" {
		fileMetadata.Parents = []string{targetFolderID}
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
