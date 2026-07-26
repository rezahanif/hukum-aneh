package main

import (
        "context"
        "log/slog"
        "os"
        "os/signal"
        "syscall"

        "github.com/rezahanif/hukum-aneh/backend/internal/ai"
        "github.com/rezahanif/hukum-aneh/backend/internal/config"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors"
        "github.com/rezahanif/hukum-aneh/backend/internal/parser"
        "github.com/rezahanif/hukum-aneh/backend/internal/repository"
        "github.com/rezahanif/hukum-aneh/backend/internal/retrieval"
        "github.com/rezahanif/hukum-aneh/backend/internal/services/imagegen"
        "github.com/rezahanif/hukum-aneh/backend/internal/services/publishing"
        "github.com/rezahanif/hukum-aneh/backend/internal/services/telegram"
        "github.com/rezahanif/hukum-aneh/backend/internal/validator"
        "github.com/rezahanif/hukum-aneh/backend/internal/workflow"
)

func main() {
        logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}))
        slog.SetDefault(logger)

        cfg, err := config.Load()
        if err != nil {
                logger.Error("config load failed", "error", err)
                os.Exit(1)
        }

        ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
        defer cancel()

        repos, err := repository.NewRepoSet(ctx, cfg)
        if err != nil {
                logger.Error("storage init failed", "error", err, "mode", cfg.StorageMode)
                os.Exit(1)
        }
        defer repos.Closer.Close()

        tgSvc := telegram.New(cfg)

        // Qdrant client: only created when storage mode is postgres or dual_write
        var qdrantClient *retrieval.QdrantClient
        if cfg.IsPostgres() {
                qdrantClient, err = retrieval.NewQdrantClient(ctx, cfg.Qdrant.Host, cfg.Qdrant.Port, cfg.Qdrant.Collection, cfg.Qdrant.APIKey, cfg.Qdrant.VectorSize, logger)
                if err != nil {
                        logger.Error("qdrant init failed", "error", err)
                        os.Exit(1)
                }
                defer qdrantClient.Close()
        }

        ret, err := retrieval.New(ctx, cfg, repos.EmbedRepo, qdrantClient)
        if err != nil {
                logger.Error("retrieval service init failed", "error", err)
                os.Exit(1)
        }

        engine := workflow.NewEngine(
                cfg,
                repos,
                connectors.NewRegistry(),
                parser.New(logger),
                ret,
                qdrantClient,
                ai.New(cfg),
                imagegen.New(cfg),
                tgSvc,
                publishing.New(cfg),
                validator.New(),
                logger,
        )

        logger.Info("telegram approval bot polling started")
        if err := tgSvc.StartPolling(ctx, func(draftID string, action string, reviewerID string) error {
                return engine.HandleApprovalAction(ctx, draftID, action, reviewerID)
        }); err != nil && ctx.Err() == nil {
                logger.Error("telegram polling stopped", "error", err)
                os.Exit(1)
        }
}
