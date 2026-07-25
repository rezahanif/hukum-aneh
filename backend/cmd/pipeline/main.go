package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rezahanif/hukum-aneh/backend/internal/config"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/peraturan"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/jdihn"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/bpk"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/mkri"
	"github.com/rezahanif/hukum-aneh/backend/internal/connectors/setneg"
	"github.com/rezahanif/hukum-aneh/backend/internal/parser"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
	"github.com/rezahanif/hukum-aneh/backend/internal/scheduler"
	"github.com/rezahanif/hukum-aneh/backend/internal/workflow"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
	"github.com/rezahanif/hukum-aneh/backend/internal/retrieval"
	"github.com/rezahanif/hukum-aneh/backend/internal/ai"
	"github.com/rezahanif/hukum-aneh/backend/internal/services/imagegen"
	"github.com/rezahanif/hukum-aneh/backend/internal/services/publishing"
	"github.com/rezahanif/hukum-aneh/backend/internal/services/telegram"
	"github.com/rezahanif/hukum-aneh/backend/internal/validator"
)

func main() {
	var (
		runOnce bool
		verbose bool
	)
	flag.BoolVar(&runOnce, "once", false, "run discovery once and exit (no scheduler loop)")
	flag.BoolVar(&verbose, "verbose", false, "enable debug logging")
	flag.Parse()

	level := slog.LevelInfo
	if verbose {
		level = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)

	cfg, err := config.Load()
	if err != nil {
		logger.Error("config load failed", "error", err)
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Firestore
	repo, err := repository.NewFirestoreRepo(ctx, cfg.Firebase.ProjectID, cfg.Firebase.CredentialsPath)
	if err != nil {
		logger.Error("firestore init failed", "error", err)
		os.Exit(1)
	}
	defer repo.Close()

	// Connectors registry
	registry := connectors.NewRegistry()

	// Python scraper bridge
	scr := scraper.New(cfg.Scraper.PythonPath, cfg.Scraper.ScriptPath, logger)

	// Register Peraturan.go.id connector
	peraturanConn := peraturan.New(scr, logger)
	registry.Register(peraturanConn.Name(), peraturanConn)

	// Register JDIHN connector
	jdihnConn := jdihn.New(scr, logger)
	registry.Register(jdihnConn.Name(), jdihnConn)

	// Register BPK connector
	bpkConn := bpk.New(logger)
	registry.Register(bpkConn.Name(), bpkConn)

	// Register MKRI connector
	mkriConn := mkri.New(scr, logger)
	registry.Register(mkriConn.Name(), mkriConn)

	// Register JDIH Setneg connector
	setnegConn := setneg.New(scr, logger)
	registry.Register(setnegConn.Name(), setnegConn)

	// Document parser
	p := parser.New(logger)

	// Retrieval (embedding & search)
	ret := retrieval.New(cfg, repo)

	// AI Agents
	aiSvc := ai.New(cfg)

	// Image Generator
	imgGen := imagegen.New(cfg)

	// Telegram Bot
	tgSvc := telegram.New(cfg)

	// Publishing Service
	pubSvc := publishing.New(cfg)

	// Image Validator
	val := validator.New()

	engine := workflow.NewEngine(cfg, repo, registry, p, ret, aiSvc, imgGen, tgSvc, pubSvc, val, logger)

	if runOnce {
		logger.Info("running discovery once")
		if err := engine.RunDiscovery(ctx); err != nil {
			logger.Error("discovery failed", "error", err)
			os.Exit(1)
		}
		return
	}

	// Scheduler
	interval, err := time.ParseDuration(cfg.Scheduler.DiscoveryInterval)
	if err != nil {
		interval = time.Hour
	}
	stuckThreshold, err := time.ParseDuration(cfg.Scheduler.StuckJobThreshold)
	if err != nil {
		stuckThreshold = 6 * time.Hour
	}

	sched := scheduler.New(
		engine.RunDiscovery,
		scheduler.WithInterval(interval),
		scheduler.WithStuckThreshold(stuckThreshold),
		scheduler.WithLogger(logger),
	)
	sched.SetStuckCheck(engine.CheckStuckJobs)

	if err := sched.Start(ctx); err != nil {
		logger.Error("scheduler start failed", "error", err)
		os.Exit(1)
	}

	// Start Telegram Polling for approvals
	go func() {
		logger.Info("starting Telegram approval bot polling")
		err := tgSvc.StartPolling(ctx, func(draftID string, action string, reviewerID string) error {
			return engine.HandleApprovalAction(ctx, draftID, action, reviewerID)
		})
		if err != nil {
			logger.Error("telegram bot polling stopped", "error", err)
		}
	}()

	// Graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	logger.Info("shutting down")
	cancel()
	sched.Stop()
}
