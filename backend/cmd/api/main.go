package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/models"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
	"github.com/rezahanif/hukum-aneh/backend/internal/retrieval"
)

type Server struct {
	cfg       *config.Config
	repos     *repository.RepoSet
	retrieval *retrieval.Service
	logger    *slog.Logger
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	ctx := context.Background()
	repos, err := repository.NewRepoSet(ctx, cfg)
	if err != nil {
		log.Fatalf("repos: %v", err)
	}
	defer repos.Closer.Close()

	var qdrantClient *retrieval.QdrantClient
	if cfg.IsPostgres() {
		var err error
		qdrantClient, err = retrieval.NewQdrantClient(ctx, cfg.Qdrant.Host, cfg.Qdrant.Port, cfg.Qdrant.Collection, cfg.Qdrant.APIKey, cfg.Qdrant.VectorSize, logger)
		if err != nil {
			logger.Warn("qdrant init failed (semantic search disabled)", "error", err)
		} else {
			defer qdrantClient.Close()
		}
	}

	ret, err := retrieval.New(ctx, cfg, repos.EmbedRepo, qdrantClient)
	if err != nil {
		log.Fatalf("retrieval: %v", err)
	}

	srv := &Server{
		cfg:       cfg,
		repos:     repos,
		retrieval: ret,
		logger:    logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api/laws", srv.handleLaws)
	mux.HandleFunc("/api/laws/", srv.handleLawDetail)
	mux.HandleFunc("/api/search", srv.handleSearch)

	addr := fmt.Sprintf(":%d", port)
	logger.Info("Starting API server", "addr", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("listen: %v", err)
	}
}

func (s *Server) handleLaws(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	status := r.URL.Query().Get("status")
	var laws []models.LawDocument
	var err error

	if status != "" {
		laws, err = s.repos.LawRepo.ListLawsByStatus(r.Context(), status)
	} else {
		laws, err = s.repos.LawRepo.ListAllLaws(r.Context())
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(laws)
}

func (s *Server) handleLawDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 3 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	id := parts[2]

	if len(parts) == 4 && parts[3] == "version" {
		v, err := s.repos.VersionRepo.GetLatestLawVersion(r.Context(), id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(v)
		return
	}

	doc, err := s.repos.LawRepo.GetLawDocument(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(doc)
}

func (s *Server) handleSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "missing query param q", http.StatusBadRequest)
		return
	}

	topN := 5
	if nStr := r.URL.Query().Get("limit"); nStr != "" {
		if val, err := strconv.Atoi(nStr); err == nil && val > 0 {
			topN = val
		}
	}

	vec, _, err := s.retrieval.GenerateEmbedding(r.Context(), query)
	if err != nil {
		http.Error(w, fmt.Sprintf("embedding generation failed: %v", err), http.StatusInternalServerError)
		return
	}

	results, err := s.retrieval.Search(r.Context(), vec, topN)
	if err != nil {
		http.Error(w, fmt.Sprintf("similarity search failed: %v", err), http.StatusInternalServerError)
		return
	}

	type SearchResponseItem struct {
		ID          string  `json:"id"`
		LawNumber   string  `json:"law_number"`
		Title       string  `json:"title"`
		Score       float32 `json:"score"`
		RawFilePath string  `json:"raw_file_path"`
	}

	var respItems []SearchResponseItem
	for _, res := range results {
		doc, err := s.repos.LawRepo.GetLawDocument(r.Context(), res.LawDocumentID)
		if err != nil {
			continue
		}
		respItems = append(respItems, SearchResponseItem{
			ID:          doc.ID,
			LawNumber:   doc.LawNumber,
			Title:       doc.Title,
			Score:       res.Score,
			RawFilePath: doc.RawFilePath,
		})
	}

	json.NewEncoder(w).Encode(respItems)
}
