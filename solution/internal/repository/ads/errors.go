package ads

import "server/pkg/errors"

var (
	ErrEndedAd             = errors.NewError("ad is ended", errors.BadRequest)
	ErrAdNotFound          = errors.NewError("ad not found", errors.NotFound)
	ErrClientNotRegistered = errors.NewError("client not registered", errors.BadRequest)
	ErrCampaignNotFound    = errors.NewError("campaign not found", errors.BadRequest)
)
