package errors

import (
	"encoding/json"
	"errors"
	stderrors "errors"
	"fmt"
	"log"
)

var (
	Is     = stderrors.Is
	As     = stderrors.As
	Unwrap = stderrors.Unwrap
)

type Error struct {
	code    Code
	message string
	cause   error
	stack   *stack
}

func (e *Error) Code() Code {
	return e.code
}

func (e *Error) Message() string {
	return e.message
}

func (e *Error) Cause() error {
	return e.cause
}

func (e *Error) Error() string {
	out := e.message
	if e.cause != nil {
		out = out + ": " + e.cause.Error()
	}
	return out
}

func (e *Error) StackTrace() string {
	return e.stack.frames().format()
}

func (e *Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}{
		Code:    e.code.String(),
		Message: e.message,
	})
}

func (e *Error) WithMessagef(message string, args ...any) *Error {
	e.message = fmt.Sprintf(message, args...)
	return e
}

func (e *Error) WithMessage(message string) *Error {
	e.message = message
	return e
}

func (e *Error) WithCode(code Code) *Error {
	e.code = code
	return e
}

func (e *Error) WithCause(cause error) *Error {
	e.cause = cause
	return e
}

func (e *Error) Unwrap() error {
	return e.cause
}

type Errors []error

var _ error = Errors{}

func (errs Errors) Error() string {
	var tmpErrs struct {
		Errors []Error `json:"errors,omitempty"`
	}

	for _, e := range errs {
		var err *Error
		ok := errors.As(e, &err)
		if !ok {
			err = UnknownError(e)
		}
		if err.code == "" {
			err.code = GeneralCode
		}

		tmpErrs.Errors = append(tmpErrs.Errors, *err)
	}

	msg, err := json.Marshal(tmpErrs)
	if err != nil {
		log.Printf("failed to marshal errors: %v", err)
		return "{}"
	}
	return string(msg)
}

func (errs Errors) Len() int {
	return len(errs)
}

func NewErrs(err error) Errors {
	return Errors{err}
}

func New(in any) *Error {
	var err error
	switch in := in.(type) {
	case error:
		err = in
	default:
		err = fmt.Errorf("%v", in)
	}

	return &Error{
		message: err.Error(),
		stack:   newStack(),
	}
}

func Wrap(err error, message string) *Error {
	if err == nil {
		return nil
	}
	e := &Error{
		cause:   err,
		message: message,
		stack:   newStack(),
	}
	return e
}

func Wrapf(err error, format string, args ...any) *Error {
	if err == nil {
		return nil
	}
	e := &Error{
		cause:   err,
		message: fmt.Sprintf(format, args...),
		stack:   newStack(),
	}
	return e
}

func Errorf(format string, args ...any) *Error {
	return &Error{
		message: fmt.Sprintf(format, args...),
		stack:   newStack(),
	}
}

func Cause(err error) error {
	for err != nil {
		var cause *Error
		ok := errors.As(err, &cause)
		if !ok {
			break
		}
		if cause.cause == nil {
			break
		}
		err = cause.cause
	}
	return err
}

func IsErr(err error, code Code) bool {
	var e *Error
	if As(err, &e) {
		return e.code == code
	}
	return false
}

func ErrCode(err error) Code {
	if err == nil {
		return ""
	}

	var e *Error
	if ok := As(err, &e); ok && e.code != "" {
		return e.code
	} else if ok && e.cause != nil {
		return ErrCode(e.cause)
	}

	return GeneralCode
}
