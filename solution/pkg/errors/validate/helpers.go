package validate

import (
	"context"

	"github.com/google/uuid"
)

func ValidateUUID(id string) Condition {
	return func(_ context.Context) error {
		_, err := uuid.Parse(id)
		if err != nil {
			return NewValidationError("invalid uuid")
		}
		return nil
	}
}

func IsValidUUID(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return NewValidationError("invalid uuid")
	}
	return nil
}
