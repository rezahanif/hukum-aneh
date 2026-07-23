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
	"github.com/rezahanif/hukum-aneh/backend/internal/parser"
	"github.com/rezahanif/hukum-aneh/backend/internal/repository"
	"github.com/rezahanif/hukum-aneh/backend/internal/scheduler"
	"github.com/rezahanif/hukum-aneh/backend/internal/workflow"
	"github.com/rezahanif/hukum-aneh/backend/pkg/scraper"
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

	// Document parser
	p := parser.New(logger)

	engine := workflow.NewEngine(cfg, repo, registry, p, logger)

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

	// Graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh

	logger.Info("shutting down")
	cancel()
	sched.Stop()
}
