package http

import (
	"net/http"

	"github.com/guodongq/uap/errors"
)

var (
	codeMap = map[errors.Code]int{
		errors.InvalidArgumentCode: http.StatusBadRequest,
		errors.UnAuthorizedCode:    http.StatusUnauthorized,
		errors.ForbiddenCode:       http.StatusForbidden,
		errors.NotFoundCode:        http.StatusNotFound,
		errors.ConflictCode:        http.StatusConflict,
		errors.TimeoutCode:         http.StatusGatewayTimeout,
		errors.UnavailableCode:     http.StatusServiceUnavailable,
		errors.NotImplementedCode:  http.StatusNotImplemented,
		errors.GeneralCode:         http.StatusInternalServerError,
		errors.InternalCode:        http.StatusInternalServerError,
	}
)

// StatusCode returns the HTTP status code for the given error.
func StatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	code := errors.ErrCode(err)
	if status, ok := codeMap[code]; ok {
		return status
	}

	return http.StatusInternalServerError
}

// FromStatus returns the error Code for the given HTTP status code.
func FromStatus(status int) errors.Code {
	switch status {
	case http.StatusBadRequest:
		return errors.InvalidArgumentCode
	case http.StatusUnauthorized:
		return errors.UnAuthorizedCode
	case http.StatusForbidden:
		return errors.ForbiddenCode
	case http.StatusNotFound:
		return errors.NotFoundCode
	case http.StatusConflict:
		return errors.ConflictCode
	case http.StatusGatewayTimeout:
		return errors.TimeoutCode
	case http.StatusServiceUnavailable:
		return errors.UnavailableCode
	case http.StatusNotImplemented:
		return errors.NotImplementedCode
	case http.StatusInternalServerError:
		return errors.InternalCode
	default:
		return errors.InternalCode
	}
}
