package logger

import (
	"log/slog"
	"os"
	"strings"
)

func Init(level, format string) {

	var levelValue slog.Level
	level = strings.ToLower(level)

	switch level {
	case "debug":
		levelValue = slog.LevelDebug
	case "info":
		levelValue = slog.LevelInfo
	case "warn":
		levelValue = slog.LevelWarn
	case "error":
		levelValue = slog.LevelError
	default:
		levelValue = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: levelValue,
	}

	var handler slog.Handler
	if strings.ToLower(format) == "text" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	slog.SetDefault(slog.New(handler))

}
