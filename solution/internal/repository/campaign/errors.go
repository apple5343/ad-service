package campaign

import "server/pkg/errors"

var (
	ErrAdvertiserConflict      = errors.NewError("advertiser conflict", errors.Conflict)
	ErrAdvertiserNotRegistered = errors.NewError("advertiser not registered", errors.BadRequest)
	ErrCampaignNotFound        = errors.NewError("campaign not found", errors.BadRequest)
	ErrPermissionDenied        = errors.NewError("permission denied", errors.Forbidden)
)
