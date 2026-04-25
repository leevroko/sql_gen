package logger

import (
	"log/slog"
	"os"
)

const (
	LevelDebug = slog.LevelDebug
	LevelInfo = slog.LevelInfo
	LevelWarn = slog.LevelWarn
	LevelError = slog.LevelError
)

func NewLogger(level slog.Level) *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level, AddSource: true}))
}

