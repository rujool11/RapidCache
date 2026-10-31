package main

import (
	"log/slog"

	"github.com/joho/godotenv"
	"github.com/rujool11/RapidCache/internal/config"
	"github.com/rujool11/RapidCache/internal/logger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Warn("No environment config found, relying on defaults")
	}

	cfg := config.LoadConfig()
	logger.Init(cfg.LogLevel, cfg.LogFormat)

	slog.Info("Initializing application...",
		"port", cfg.Port,
		"log_level", cfg.LogLevel,
		"log_format", cfg.LogFormat)

}
