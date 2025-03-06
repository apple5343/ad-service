package validate

import (
	"encoding/json"
	"errors"
)

type ValidationError struct {
	messages []string
}

func NewValidationError(messages ...string) *ValidationError {
	return &ValidationError{messages}
}

func (v *ValidationError) Add(message string) {
	v.messages = append(v.messages, message)
}

func (v *ValidationError) Messages() []string {
	return v.messages
}

func (v *ValidationError) Error() string {
	data, err := json.Marshal(v.messages)
	if err != nil {
		return err.Error()
	}

	return string(data)
}

func IsValidationError(err error) bool {
	var ve *ValidationError
	return errors.As(err, &ve)
}
