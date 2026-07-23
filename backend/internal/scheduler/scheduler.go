package scheduler

import (
	"context"
	"log/slog"
	"sync"
	"time"
)

// Scheduler triggers discovery runs on interval and re-drives stalled stages.
// Pure deterministic — no AI, no reasoning. Spec §5.1.
type Scheduler struct {
	interval       time.Duration
	stuckThreshold time.Duration
	discoveryFn    func(ctx context.Context) error
	stuckCheckFn   func(ctx context.Context) error
	logger         *slog.Logger
	cancel         context.CancelFunc
	wg             sync.WaitGroup
	mu             sync.Mutex
	running        bool
}

type Option func(*Scheduler)

func WithInterval(d time.Duration) Option {
	return func(s *Scheduler) { s.interval = d }
}

func WithStuckThreshold(d time.Duration) Option {
	return func(s *Scheduler) { s.stuckThreshold = d }
}

func WithLogger(l *slog.Logger) Option {
	return func(s *Scheduler) { s.logger = l }
}

// New creates a Scheduler. discoveryFn is called on each tick.
// stuckCheckFn is optional — called after each discovery tick.
func New(discoveryFn func(ctx context.Context) error, opts ...Option) *Scheduler {
	s := &Scheduler{
		interval:       time.Hour,
		stuckThreshold: 6 * time.Hour,
		discoveryFn:    discoveryFn,
		logger:         slog.Default(),
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (s *Scheduler) SetStuckCheck(fn func(ctx context.Context) error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stuckCheckFn = fn
}

// Start begins the scheduler loop. Blocks until Stop is called.
func (s *Scheduler) Start(parentCtx context.Context) error {
	s.mu.Lock()
	if s.running {
		s.mu.Unlock()
		return nil
	}
	ctx, cancel := context.WithCancel(parentCtx)
	s.cancel = cancel
	s.running = true
	s.mu.Unlock()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		ticker := time.NewTicker(s.interval)
		defer ticker.Stop()

		s.logger.Info("scheduler started", "interval", s.interval)

		// Run immediately on start.
		s.tick(ctx)

		for {
			select {
			case <-ctx.Done():
				s.logger.Info("scheduler stopped")
				return
			case <-ticker.C:
				s.tick(ctx)
			}
		}
	}()
	return nil
}

func (s *Scheduler) tick(ctx context.Context) {
	if err := s.discoveryFn(ctx); err != nil {
		s.logger.Error("discovery tick failed", "error", err)
	}
	if s.stuckCheckFn != nil {
		if err := s.stuckCheckFn(ctx); err != nil {
			s.logger.Error("stuck check failed", "error", err)
		}
	}
}

// Stop gracefully shuts down the scheduler.
func (s *Scheduler) Stop() {
	s.mu.Lock()
	if !s.running {
		s.mu.Unlock()
		return
	}
	s.running = false
	if s.cancel != nil {
		s.cancel()
	}
	s.mu.Unlock()
	s.wg.Wait()
}
