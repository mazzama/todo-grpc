package custom_error

import "errors"

var (
	ErrInvalidStatus = errors.New("invalid status value")
	ErrEmptyField    = errors.New("field should not empty")
)
