package domain

import "errors"

var (
	ErrInternalServerError = errors.New("Internal server error")
	ErrBadParamInput       = errors.New("Invalid request body")
)
