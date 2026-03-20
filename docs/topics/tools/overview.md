---
sidebar_position: 1
---

# Tools

The `tools` package contains general-purpose utility packages organized into four categories.

## Concurrency

Utilities for concurrent programming:

- **executor** — Goroutine lifecycle management
- **pool** — Worker pool implementation
- **semaphore** — Counting semaphore

## Container

Generic data structure utilities:

- **maps** — Concurrent map, ordered map, map helpers
- **set** — Generic set implementation
- **slice** — Slice manipulation (filter, map, shuffle, etc.)

## Lang

Go language extensions:

- **ptr** — Pointer helpers (`ToPtr`, `FromPtr`)
- **clone** — Deep copy
- **ctxutil** — Type-safe context get/set with generics
- **encodingx** — Base64 utilities
- **mathx** — Math helpers
- **randx** — Random generation utilities

## Sys

System-level utilities:

- **env** — Environment variable helpers
- **shell** — Command execution
- **fswatcher** — Filesystem watcher
- **retry** — Retry with backoff
