package campaign

import (
	"context"
	"server/pkg/errors/validate"
)

func (s *campaignService) Delete(ctx context.Context, advertiserID, campaignID string) error {
	if err := validate.IsValidUUID(advertiserID); err != nil {
		return validate.NewValidationError("invalid advertiser id")
	}
	if err := validate.IsValidUUID(campaignID); err != nil {
		return validate.NewValidationError("invalid campaign id")
	}
	return s.repository.Delete(ctx, advertiserID, campaignID)
}
