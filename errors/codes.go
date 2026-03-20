package errors

type Code string

const (
	GeneralCode  Code = "UNKNOWN"
	InternalCode Code = "INTERNAL"

	InvalidArgumentCode Code = "INVALID_ARGUMENT"

	UnAuthorizedCode Code = "UNAUTHORIZED"
	ForbiddenCode    Code = "FORBIDDEN"

	NotFoundCode Code = "NOT_FOUND"
	ConflictCode Code = "CONFLICT"

	TimeoutCode        Code = "TIMEOUT"
	UnavailableCode    Code = "UNAVAILABLE"
	NotImplementedCode Code = "NOT_IMPLEMENTED"
)

var validCodes = map[Code]struct{}{
	GeneralCode:         {},
	InternalCode:        {},
	InvalidArgumentCode: {},
	UnAuthorizedCode:    {},
	ForbiddenCode:       {},
	NotFoundCode:        {},
	ConflictCode:        {},
	TimeoutCode:         {},
	UnavailableCode:     {},
	NotImplementedCode:  {},
}

func (c Code) String() string {
	return string(c)
}

func (c Code) IsValid() bool {
	_, ok := validCodes[c]
	return ok
}

func InternalError(err error) *Error {
	return New("internal error").WithCode(InternalCode).WithCause(err)
}

func InvalidArgumentError(err error) *Error {
	return New("invalid argument").WithCode(InvalidArgumentCode).WithCause(err)
}

func UnAuthorizedError(err error) *Error {
	return New("unauthorized").WithCode(UnAuthorizedCode).WithCause(err)
}

func ForbiddenError(err error) *Error {
	return New("forbidden").WithCode(ForbiddenCode).WithCause(err)
}

func NotFoundError(err error) *Error {
	return New("resource not found").WithCode(NotFoundCode).WithCause(err)
}

func ConflictError(err error) *Error {
	return New("conflict").WithCode(ConflictCode).WithCause(err)
}

func TimeoutError(err error) *Error {
	return New("timeout").WithCode(TimeoutCode).WithCause(err)
}

func UnavailableError(err error) *Error {
	return New("unavailable").WithCode(UnavailableCode).WithCause(err)
}

func NotImplementedError(err error) *Error {
	return New("not implemented").WithCode(NotImplementedCode).WithCause(err)
}

func UnknownError(err error) *Error {
	return New("unknown").WithCode(GeneralCode).WithCause(err)
}

func IsUnknownError(err error) bool {
	return IsErr(err, GeneralCode)
}

func IsInternalError(err error) bool {
	return IsErr(err, InternalCode)
}

func IsInvalidArgumentError(err error) bool {
	return IsErr(err, InvalidArgumentCode)
}

func IsUnAuthorizedError(err error) bool {
	return IsErr(err, UnAuthorizedCode)
}

func IsForbiddenError(err error) bool {
	return IsErr(err, ForbiddenCode)
}

func IsNotFoundError(err error) bool {
	return IsErr(err, NotFoundCode)
}

func IsConflictError(err error) bool {
	return IsErr(err, ConflictCode)
}

func IsTimeoutError(err error) bool {
	return IsErr(err, TimeoutCode)
}

func IsUnavailableError(err error) bool {
	return IsErr(err, UnavailableCode)
}

func IsNotImplementedError(err error) bool {
	return IsErr(err, NotImplementedCode)
}
