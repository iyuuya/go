package logger

import (
	"io"
	"log/slog"
)

const (
	levelDebug = "DEBUG"
	levelInfo  = "INFO"
	levelWarn  = "WARN"
	levelError = "ERROR"
)

var initialized bool

func Init(w io.Writer, level string) {
	if initialized {
		slog.Warn("logger already initialized")
		return
	}

	var l slog.Level
	switch level {
	case levelDebug:
		l = slog.LevelDebug
	case levelInfo:
		l = slog.LevelInfo
	case levelWarn:
		l = slog.LevelWarn
	case levelError:
		l = slog.LevelError
	default:
		l = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{Level: l}))
	slog.SetDefault(logger)
	initialized = true
}
