# SDC logger

![sdc logo](https://docs.sdcio.dev/assets/logos/SDC-transparent-withname-100x133.png)


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

### Logger options
To set custom options, pass the wanted log level and ReplaceAttr function in the `options` parameter when creating the logger:

**NOTE: the conversion between slog and logr levels results in negating the level, slog levels are more verbose the more negative the number, so here we use the library's constants negated to set the slog level, which lets us have the expected behaviour when using these constants in code.**

```go
package main

import (
	"context"
	"log/slog"

	"github.com/go-logr/logr"
	logf "github.com/sdcio/logger"
)

func main() {
	slogOpts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: logf.ReplaceTimeAttr,
	}

	if debug {
		slogOpts.Level = slog.Level(-logf.VDebug)
	}
	if trace {
		slogOpts.Level = slog.Level(-logf.VTrace)
	}

	log := logr.FromSlogHandler(slog.NewJSONHandler(os.Stdout, slogOpts))
	logf.SetDefaultLogger(log)
}
```


## Join us

Have questions, ideas, bug reports or just want to chat? Come join [our discord server](https://discord.com/channels/1240272304294985800/1311031796372344894).

## License and Code of Conduct

Code is under the [Apache License 2.0](LICENSE), documentation is [CC BY 4.0](LICENSE-documentation).

The SDC project is following the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/main/code-of-conduct.md). More information and links about the CNCF Code of Conduct are [here](https://www.cncf.io/conduct/).
