package helpers

import "errors"

var (
	ErrBadRequest      = errors.New("bad request")
	ErrInternalFaliure = errors.New("internal faliure")
	ErrNotFound        = errors.New("not found")
	ErrInvalidPayload  = errors.New("invalid payload")
)
