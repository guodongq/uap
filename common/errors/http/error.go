package http

import (
	"net/http"

	errors2 "github.com/guodongq/uap/common/errors"
)

var (
	codeMap = map[errors2.Code]int{
		errors2.InvalidArgumentCode: http.StatusBadRequest,
		errors2.UnAuthorizedCode:    http.StatusUnauthorized,
		errors2.ForbiddenCode:       http.StatusForbidden,
		errors2.NotFoundCode:        http.StatusNotFound,
		errors2.ConflictCode:        http.StatusConflict,
		errors2.TimeoutCode:         http.StatusGatewayTimeout,
		errors2.UnavailableCode:     http.StatusServiceUnavailable,
		errors2.NotImplementedCode:  http.StatusNotImplemented,
		errors2.GeneralCode:         http.StatusInternalServerError,
		errors2.InternalCode:        http.StatusInternalServerError,
	}
)

// StatusCode returns the HTTP status code for the given error.
func StatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	code := errors2.ErrCode(err)
	if status, ok := codeMap[code]; ok {
		return status
	}

	return http.StatusInternalServerError
}

// FromStatus returns the error Code for the given HTTP status code.
func FromStatus(status int) errors2.Code {
	switch status {
	case http.StatusBadRequest:
		return errors2.InvalidArgumentCode
	case http.StatusUnauthorized:
		return errors2.UnAuthorizedCode
	case http.StatusForbidden:
		return errors2.ForbiddenCode
	case http.StatusNotFound:
		return errors2.NotFoundCode
	case http.StatusConflict:
		return errors2.ConflictCode
	case http.StatusGatewayTimeout:
		return errors2.TimeoutCode
	case http.StatusServiceUnavailable:
		return errors2.UnavailableCode
	case http.StatusNotImplemented:
		return errors2.NotImplementedCode
	case http.StatusInternalServerError:
		return errors2.InternalCode
	default:
		return errors2.InternalCode
	}
}
