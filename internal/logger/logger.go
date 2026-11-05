package logger

import (
	"log/slog"
	"os"
	"strings"
)

func Init(level, format string) {

	var levelValue slog.Level
	switch level {
	case "DEBUG":
		levelValue = slog.LevelDebug
	case "INFO":
		levelValue = slog.LevelInfo
	case "WARN":
		levelValue = slog.LevelWarn
	case "ERROR":
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
