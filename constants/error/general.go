package error

import "errors"

const (
	Success = "success"
	Error   = "error"
)

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrSQLError            = errors.New("Database server failed to execute query")
	ErrToManyRequests      = errors.New("Too many requests")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrInvalidToken        = errors.New("Invalid token")
	ErrForbidden           = errors.New("Forbidden")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrToManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}
