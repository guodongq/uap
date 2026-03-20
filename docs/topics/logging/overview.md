---
sidebar_position: 1
---

# Logging

UAP provides a logging abstraction layer that decouples your application from specific logging implementations.

## Supported Backends

- **Logrus** — `github.com/guodongq/uap/log/logrus`
- **Zap** — `github.com/guodongq/uap/log/zap`

## Basic Usage

```go
import "github.com/guodongq/uap/log"

log.Info("server started")
log.WithFields(log.Fields{"port": 8080}).Info("listening")
log.Error("something went wrong")
```

## Context-aware Logging

```go
ctx := log.WithLogger(ctx, logger)
log.FromContext(ctx).Info("request processed")
```
