package grpc

import (
	errors2 "github.com/guodongq/uap/common/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	codeMap = map[errors2.Code]codes.Code{
		errors2.InvalidArgumentCode: codes.InvalidArgument,
		errors2.UnAuthorizedCode:    codes.Unauthenticated,
		errors2.ForbiddenCode:       codes.PermissionDenied,
		errors2.NotFoundCode:        codes.NotFound,
		errors2.ConflictCode:        codes.Aborted,
		errors2.TimeoutCode:         codes.DeadlineExceeded,
		errors2.UnavailableCode:     codes.Unavailable,
		errors2.NotImplementedCode:  codes.Unimplemented,
		errors2.GeneralCode:         codes.Unknown,
		errors2.InternalCode:        codes.Internal,
	}
)

// GrpcCode returns the gRPC status code for the given error.
func GrpcCode(err error) codes.Code {
	if err == nil {
		return codes.OK
	}

	// Try to get gRPC status directly if it's already a gRPC error
	if s, ok := status.FromError(err); ok {
		return s.Code()
	}

	code := errors2.ErrCode(err)
	if c, ok := codeMap[code]; ok {
		return c
	}

	return codes.Unknown
}

// Status returns the gRPC status for the given error.
func Status(err error) *status.Status {
	if err == nil {
		return status.New(codes.OK, "")
	}

	// Try to get gRPC status directly if it's already a gRPC error
	if s, ok := status.FromError(err); ok {
		return s
	}

	return status.New(GrpcCode(err), err.Error())
}

// FromCode maps a standard gRPC status code to an error Code.
func FromCode(code codes.Code) errors2.Code {
	switch code {
	case codes.InvalidArgument:
		return errors2.InvalidArgumentCode
	case codes.Unauthenticated:
		return errors2.UnAuthorizedCode
	case codes.PermissionDenied:
		return errors2.ForbiddenCode
	case codes.NotFound:
		return errors2.NotFoundCode
	case codes.Aborted:
		return errors2.ConflictCode
	case codes.DeadlineExceeded:
		return errors2.TimeoutCode
	case codes.Unavailable:
		return errors2.UnavailableCode
	case codes.Unimplemented:
		return errors2.NotImplementedCode
	case codes.Internal:
		return errors2.InternalCode
	case codes.Unknown:
		return errors2.GeneralCode
	default:
		return errors2.GeneralCode
	}
}
