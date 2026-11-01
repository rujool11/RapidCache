package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/rujool11/RapidCache/internal/config"
	"github.com/rujool11/RapidCache/internal/logger"
	httpserver "github.com/rujool11/RapidCache/internal/server/http"
	respserver "github.com/rujool11/RapidCache/internal/server/resp"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Fatal %v \n", err)
		os.Exit(1)
	}
}

func run() error {
	if err := godotenv.Load(); err != nil {
		slog.Warn("No environment config found, relying on defaults")
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	logger.Init(cfg.LogLevel, cfg.LogFormat)

	slog.Info("Initializing application...",
		"log_level", cfg.LogLevel,
		"log_format", cfg.LogFormat,
		"resp_port", cfg.RESPPort,
		"http_enabled", cfg.HTTPEnabled,
		"http_port", cfg.HTTPPort,
	)

	// context is a value that carries a stop signal
	// returns a context that is cancelled when cancellation signal is received (Ctrl C/ SIGTERM)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	// stop deferred for graceful shutdown
	defer stop()

	// errgroup provides easy way to add waitgroup/block (g.Wait)
	// also helps error handling (stores first non-nil error to return to wait)

	// run both servers concurrently; if either returns an error, ctx is
	// cancelled so the other shuts down too
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error { return respserver.Run(ctx, cfg) })
	if cfg.HTTPEnabled == "true" {
		g.Go(func() error { return httpserver.Run(ctx, cfg) })
	}

	// blocks until every g.Go goroutine has finished, and returns first non-nil error/ nil
	return g.Wait()
}
