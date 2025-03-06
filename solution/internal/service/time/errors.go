package time

import (
	"server/pkg/errors/validate"
)

var (
	ErrInvalidDay = validate.NewValidationError("invalid day")
)
