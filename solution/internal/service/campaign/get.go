package campaign

import (
	"context"
	"server/internal/model"
	"server/pkg/errors/validate"
)

func (s *campaignService) Get(ctx context.Context, advertiserID, campaignID string) (*model.Campaign, error) {
	if err := validate.IsValidUUID(advertiserID); err != nil {
		return nil, validate.NewValidationError("invalid advertiser id")
	}
	if err := validate.IsValidUUID(campaignID); err != nil {
		return nil, validate.NewValidationError("invalid campaign id")
	}
	return s.repository.Get(ctx, advertiserID, campaignID)
}

func (s *campaignService) List(ctx context.Context, advertiserID string, page, size int) ([]*model.Campaign, error) {
	if err := validate.IsValidUUID(advertiserID); err != nil {
		return nil, validate.NewValidationError("invalid advertiser id")
	}
	return s.repository.List(ctx, advertiserID, page, size)
}
