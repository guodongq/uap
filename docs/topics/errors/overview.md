---
sidebar_position: 1
---

# Error Handling

UAP provides a unified error handling package with built-in support for error codes, stack traces, and status mapping for gRPC and HTTP.

## Error Codes

```go
import "github.com/guodongq/uap/errors"

err := errors.New(errors.NotFound, "user not found")
err := errors.New(errors.InvalidArgument, "email is required")
```

## gRPC Status Mapping

```go
import "github.com/guodongq/uap/errors/grpcerr"

// Convert UAP error to gRPC status
st := grpcerr.ToStatus(err)
```

## HTTP Status Mapping

```go
import "github.com/guodongq/uap/errors/httperr"

// Convert UAP error to HTTP status code
code := httperr.ToHTTPStatus(err)
```
