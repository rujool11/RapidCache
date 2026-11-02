package respserver

import (
	"context"
	"fmt"

	"github.com/rujool11/RapidCache/internal/config"
)

func Run(ctx context.Context, cfg *config.Config) error {
	fmt.Println("Running RESP server...")

	<-ctx.Done()
	return teardown()
}

func teardown() error {
	return nil
}
