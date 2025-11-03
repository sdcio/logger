package logger

import (
	"log/slog"
	"time"
)

func ReplaceTimeAttr(_ []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey && a.Value.Kind() == slog.KindTime {
		t := a.Value.Time()
		a.Value = slog.StringValue(t.UTC().Format(time.RFC3339))
	}
	return a
}
