package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/go-logr/logr"
)

const (
	VTrace = -8
	VDebug = int(slog.LevelDebug)
	VWarn  = int(slog.LevelWarn)
	VError = int(slog.LevelError)
)

var fallbackLogger logr.Logger
var DefaultLogger logr.Logger

func init() {
	DefaultLogger = logr.FromSlogHandler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(VTrace),
		ReplaceAttr: ReplaceTimeAttr,
	}))
	fallbackLogger = DefaultLogger.WithName("fallback-logger")
}

func SetDefaultLogger(logger logr.Logger) {
	DefaultLogger = logger
}

func FromContext(ctx context.Context) logr.Logger {
	log := fallbackLogger
	if cLog, err := logr.FromContext(ctx); err == nil {
		log = cLog
	}
	return log
}

func IntoContext(ctx context.Context, logger logr.Logger) context.Context {
	return logr.NewContext(ctx, logger)
}
