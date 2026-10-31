package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/rujool11/RapidCache/internal/config"
	"github.com/rujool11/RapidCache/internal/logger"
	httpserver "github.com/rujool11/RapidCache/internal/server/http"
	respserver "github.com/rujool11/RapidCache/internal/server/resp"
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

	httpserver.RunHTTPServer()
	respserver.RunRESPServer()
	return nil
}
