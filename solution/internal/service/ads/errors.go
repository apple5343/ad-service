package ads

import "server/pkg/errors"

var (
	ErrNotShown = errors.NewError("ad not shown", errors.Forbidden)
)
