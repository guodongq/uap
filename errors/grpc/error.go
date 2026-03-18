package grpc

import (
	"github.com/guodongq/uap/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	codeMap = map[errors.Code]codes.Code{
		errors.InvalidArgumentCode: codes.InvalidArgument,
		errors.UnAuthorizedCode:    codes.Unauthenticated,
		errors.ForbiddenCode:       codes.PermissionDenied,
		errors.NotFoundCode:        codes.NotFound,
		errors.ConflictCode:        codes.Aborted,
		errors.TimeoutCode:         codes.DeadlineExceeded,
		errors.UnavailableCode:     codes.Unavailable,
		errors.NotImplementedCode:  codes.Unimplemented,
		errors.GeneralCode:         codes.Unknown,
		errors.InternalCode:        codes.Internal,
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

	code := errors.ErrCode(err)
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
func FromCode(code codes.Code) errors.Code {
	switch code {
	case codes.InvalidArgument:
		return errors.InvalidArgumentCode
	case codes.Unauthenticated:
		return errors.UnAuthorizedCode
	case codes.PermissionDenied:
		return errors.ForbiddenCode
	case codes.NotFound:
		return errors.NotFoundCode
	case codes.Aborted:
		return errors.ConflictCode
	case codes.DeadlineExceeded:
		return errors.TimeoutCode
	case codes.Unavailable:
		return errors.UnavailableCode
	case codes.Unimplemented:
		return errors.NotImplementedCode
	case codes.Internal:
		return errors.InternalCode
	case codes.Unknown:
		return errors.GeneralCode
	default:
		return errors.GeneralCode
	}
}
