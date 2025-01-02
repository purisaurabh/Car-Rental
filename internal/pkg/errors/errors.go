package errors

import "errors"

// User Registration and Login Errors
var (
	ErrParameterMissing = errors.New("parameter missing ")
	ErrInvalidFormat    = errors.New("invalid request format")
	ErrInternalServer   = errors.New("internal server error")
	ErrTokenEmpty             = errors.New("token string is empty")
	ErrSigningMethod          = errors.New("unexpected signing method")
	ErrInvalidToken           = errors.New("invalid token")
)
