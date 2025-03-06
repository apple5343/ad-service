package client

import (
	"server/pkg/errors"
)

var (
	ErrClientAlreadyExists = errors.NewError("client already exists", errors.Conflict)
	ErrClientNotFound      = errors.NewError("client not found", errors.NotFound)
)
