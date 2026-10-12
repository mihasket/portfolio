package logger

import (
	"log/slog"
	"os"
)

func New(env string) *slog.Logger {
	opts := &slog.HandlerOptions{Level: slog.LevelInfo}

	var handler slog.Handler
	switch env {
	case "production":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	default:
		opts.Level = slog.LevelDebug
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}
