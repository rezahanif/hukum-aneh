package main

import (
        "context"
        "flag"
        "log/slog"
        "os"
        "os/signal"
        "syscall"
        "time"

        "github.com/rezahanif/hukum-aneh/backend/internal/ai"
        "github.com/rezahanif/hukum-aneh/backend/internal/config"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/bkn"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/bpk"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/dpr"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/jdihn"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/kemendag"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/kemenkeu"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/kemnaker"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/komdigi"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/kpu"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/lkpp"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/ma"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/mkri"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/peraturan"
        "github.com/rezahanif/hukum-aneh/backend/internal/connectors/setneg"
        "github.com/rezahanif/hukum-aneh/backend/internal/parser"
        "github.com/rezahanif/hukum-aneh/backend/internal/repository"
        "github.com/rezahanif/hukum-aneh/backend/internal/retrieval"
        "github.com/rezahanif/hukum-aneh/backend/internal/scheduler"
        "github.com/rezahanif/hukum-aneh/backend/internal/services/imagegen"
        "github.com/rezahanif/hukum-aneh/backend/internal/services/publishing"
        "github.com/rezahanif/hukum-aneh/backend/internal/services/telegram"
        "github.com/rezahanif/hukum-aneh/backend/internal/validator"
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

        // Storage (firestore | postgres | dual_write based on STORAGE_MODE)
        repos, err := repository.NewRepoSet(ctx, cfg)
        if err != nil {
                logger.Error("storage init failed", "error", err, "mode", cfg.StorageMode)
                os.Exit(1)
        }
        defer repos.Closer.Close()

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
        bpkConn := bpk.New(scr, logger)
        registry.Register(bpkConn.Name(), bpkConn)

        // Register MKRI connector
        mkriConn := mkri.New(scr, logger)
        registry.Register(mkriConn.Name(), mkriConn)

        // Register JDIH Setneg connector
        setnegConn := setneg.New(scr, logger)
        registry.Register(setnegConn.Name(), setnegConn)

        // Register JDIH Kemenkeu connector (Permen/Kepmen Keuangan)
        kemenkeuConn := kemenkeu.New(scr, logger)
        registry.Register(kemenkeuConn.Name(), kemenkeuConn)

        // Register Mahkamah Agung putusan connector
        maConn := ma.New(logger)
        registry.Register(maConn.Name(), maConn)

        // Register JDIH Kemnaker connector
        kemnakerConn := kemnaker.New(scr, logger)
        registry.Register(kemnakerConn.Name(), kemnakerConn)

        // Register JDIH Kemendag connector
        kemendagConn := kemendag.New(scr, logger)
        registry.Register(kemendagConn.Name(), kemendagConn)

        // Register JDIH Komdigi connector
        komdigiConn := komdigi.New(scr, logger)
        registry.Register(komdigiConn.Name(), komdigiConn)

        // Register JDIH KPU connector
        kpuConn := kpu.New(scr, logger)
        registry.Register(kpuConn.Name(), kpuConn)

        // Register JDIH BKN connector
        bknConn := bkn.New(scr, logger)
        registry.Register(bknConn.Name(), bknConn)

        // Register JDIH LKPP connector
        lkppConn := lkpp.New(scr, logger)
        registry.Register(lkppConn.Name(), lkppConn)

        // Register JDIH DPR RI connector
        dprConn := dpr.New(scr, logger)
        registry.Register(dprConn.Name(), dprConn)

        // Document parser
        p := parser.New(logger)

        // Retrieval (embedding & search)
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

        engine := workflow.NewEngine(cfg, repos, registry, p, ret, qdrantClient, aiSvc, imgGen, tgSvc, pubSvc, val, logger)

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
