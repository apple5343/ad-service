package advertiser

import "server/pkg/errors"

var (
	ErrInvalidScore = errors.NewError("invalid score", errors.BadRequest)
)
