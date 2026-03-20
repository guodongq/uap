---
sidebar_position: 1
---

# Introduction

**UAP** (Unified Application Platform) is a Go utility library that provides a collection of well-organized, reusable packages for building Go applications.

## Installation

```bash
go get github.com/guodongq/uap
```

## Package Overview

| Package | Description |
|---------|-------------|
| `errors` | Unified error handling with gRPC/HTTP status mapping |
| `log` | Structured logging abstraction (logrus, zap adapters) |
| `model` | Common domain model types (pagination, filtering, metadata) |
| `auth` | Authentication scheme abstractions |
| `adapters` | External system connectors (fx integration) |
| `tools` | General-purpose utilities (concurrency, containers, lang, sys) |
| `version` | Build version injection |

## Quick Example

```go
package main

import (
    "github.com/guodongq/uap/errors"
    "github.com/guodongq/uap/log"
)

func main() {
    err := errors.New(errors.NotFound, "resource not found")
    log.Error(err.Error())
}
```
