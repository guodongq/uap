# UAP

Go SDK for **Unify Application Platform** — a collection of well-designed, reusable Go packages for building backend services.

[![Go Reference](https://pkg.go.dev/badge/github.com/guodongq/uap.svg)](https://pkg.go.dev/github.com/guodongq/uap)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

## Installation

```bash
go get github.com/guodongq/uap
```

## Packages

| Package | Description |
|---|---|
| [`errors`](errors/) | Structured error handling with error codes, stack traces, and HTTP/gRPC mapping |
| [`log`](log/) | Unified logging interface with Logrus and Zap adapters |
| [`model`](model/) | Domain model utilities: pagination, filtering, sorting, nullable values, audit metadata |
| [`auth`](auth/) | Authentication scheme definitions and properties |
| [`version`](version/) | Build-time version information injection |
| [`adapters/fx`](adapters/fx/) | [Uber Fx](https://github.com/uber-go/fx) framework integration |
| [`tools/concurrency`](tools/concurrency/) | Executor, pool, and semaphore for concurrent task management |
| [`tools/container`](tools/container/) | Generic maps, sets, and slice utilities |
| [`tools/lang`](tools/lang/) | Deep copy, context helpers, pointer conversions, math, random generation |
| [`tools/sys`](tools/sys/) | Environment variables, file system watcher, retry, shell utilities |

## Quick Start

### Error Handling

```go
import "github.com/guodongq/uap/errors"

err := errors.NotFoundError(fmt.Errorf("user %s not found", userID))
if errors.IsNotFoundError(err) {
    // handle not found
}
```

### Logging

```go
import "github.com/guodongq/uap/log"

log.Info("server started", "port", 8080)
log.WithField("request_id", reqID).Error("request failed")

// Context-aware logging
logger := log.G(ctx)
logger.Info("processing request")
```

### Pagination & Querying

```go
import "github.com/guodongq/uap/model"

query := model.NewQuery().
    SetPage(1, 20).
    AddSortDesc("created_at")
```

### Generic Set

```go
import "github.com/guodongq/uap/tools/container/set"

s := set.New("a", "b", "c")
s.Add("d")
s.Exists("a") // true

other := set.New("b", "c", "e")
s.Intersection(other) // {"b", "c"}
```

## Documentation

Full documentation is available at [https://guodongq.github.io/uap/](https://guodongq.github.io/uap/).

## License

[Apache License 2.0](LICENSE)
