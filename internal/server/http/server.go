package httpserver

import (
	"context"
	"fmt"

	"github.com/rujool11/RapidCache/internal/config"
)

func Run(ctx context.Context, cfg *config.Config) error {
	fmt.Println("Running HTTP server...")
	return nil
}
