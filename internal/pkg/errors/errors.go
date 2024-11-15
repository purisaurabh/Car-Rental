package errors

import "errors"

// User Registration and Login Errors
var (
	ErrParameterMissing = errors.New("parameter missing ")
	ErrInvalidFormat    = errors.New("invalid request format")
	ErrInternalServer   = errors.New("internal server error")
)
