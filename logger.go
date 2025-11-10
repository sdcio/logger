package logger

import (
	"context"
	"log/slog"
	"os"

	"github.com/go-logr/logr"
)

// slog uses negative levels to increase verbosity, but logr uses positive
// There is a mapping such that calling log.V(x) will translate this to call
// slog's handler with Level(-x). We therefore need to use increasingly positive
// values to imply increased log level
const (
	VDebug = -int(slog.LevelDebug)
	VTrace = 8
)

var fallbackLogger logr.Logger
var DefaultLogger logr.Logger

func init() {
	DefaultLogger = logr.FromSlogHandler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.Level(VTrace),
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
