package logger

import (
	"log/slog"
	"os"
)

var logger *slog.Logger

func init() {
	var programLevel = new(slog.LevelVar)
	if os.Getenv("SERVER_ENV") == "dev" {
		programLevel.Set(slog.LevelDebug)
	}
	h := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     programLevel,
	})
	logger = slog.New(h)
}

func Logger() *slog.Logger {
	return logger
}
