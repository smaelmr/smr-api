package filter

import "errors"

var (
	ErrInvalidDateRange = errors.New("data final deve ser maior ou igual a data inicial")
)
