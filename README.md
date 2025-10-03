# logger

## Usage Example

This package exposes functions to help with logging using the `logr.Logger` interface, mainly regarding context-based passing of the logger instance.

Any logging implementation which implements `logr.Logger` can be used, and there are many implementations which provide the `logr.Logger` interface (`zapr`, `logr`, ...).

It also contains a few preset log levels corresponding to levels provided by slog.

```go
package main

import (
	"context"
	"log/slog"

	"github.com/go-logr/logr"
	logf "github.com/sdcio/logger"
)

func main() {
	log := logr.FromSlogHandler(slog.NewJSONHandler(os.Stdout, nil))
	logf.SetDefaultLogger(log)

	ctx := logf.IntoContext(context.Background(), log)

	log.Info("first log message", "key", "value")

    RunContext(ctx)

    log = log.WithValues("additionalkey", "key2")
    ctx = logf.IntoContext(ctx, log)
    RunContext(ctx)
}

func RunContext(ctx context.Context) {
    log := logf.FromContext(ctx)
    log.Info("log from context logger")
}
```