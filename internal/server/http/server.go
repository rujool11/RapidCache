package httpserver

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rujool11/RapidCache/internal/config"
)

func Run(ctx context.Context, cfg *config.Config) error {
	mux := http.NewServeMux()
	addRoutes(mux)

	srv := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: mux,
	}

	// goroutine - async background process
	// channel - communication for goroutines
	// buffered error channel of capacity 1 (unbuffered blocks until another goroutine is ready to receive)
	errChan := make(chan error, 1)
	fmt.Println("Running HTTP server")

	go func() {
		// dont treat shutdown signal as a crash
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		return srv.Shutdown(context.Background())
	case err := <-errChan:
		return err
	}
}
