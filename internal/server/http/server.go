package httpserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/rujool11/RapidCache/internal/config"
)

type Server struct {
	cfg    *config.Config
	logger *slog.Logger
}

func Run(ctx context.Context, cfg *config.Config) error {
	logger := slog.With("transport", "http", "port", cfg.HTTPPort)

	srv := &Server{
		cfg:    cfg,
		logger: logger,
	}

	return srv.start(ctx)
}

func (s *Server) start(ctx context.Context) error {
	mux := http.NewServeMux()
	s.addRoutes(mux)

	srv := &http.Server{
		Addr:    ":" + s.cfg.HTTPPort,
		Handler: mux,
	}

	// goroutine - async background process
	// channel - communication for goroutines
	// buffered error channel of capacity 1 (unbuffered blocks until another goroutine is ready to receive)
	errChan := make(chan error, 1)

	go func() {
		// dont treat shutdown signal as a crash
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	fmt.Println("Running HTTP server...")

	select {
	case <-ctx.Done():
		fmt.Println("Closing HTTP server...")
		return srv.Shutdown(context.Background())
	case err := <-errChan:
		return err
	}
}
