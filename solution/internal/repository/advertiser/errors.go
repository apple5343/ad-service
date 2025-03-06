package advertiser

import (
	"server/pkg/errors"
)

var (
	ErrAdvertiserNotFound         = errors.NewError("advertiser not found", errors.BadRequest)
	ErrAdvertiserOrClientNotFound = errors.NewError("advertiser or client not found", errors.BadRequest)
	ErrAdvertiserAlreadyExists    = errors.NewError("advertiser already exists", errors.Conflict)
)
