package logger

import (
	"log/slog"
	"os"
)

func New(level string) *slog.Logger {
	opts := &slog.HandlerOptions{Level: ParseLevel(level)}
	return slog.New(slog.NewJSONHandler(os.Stdout, opts))
}

func ParseLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
