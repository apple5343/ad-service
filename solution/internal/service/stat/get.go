package stat

import (
	"context"
	"server/internal/model"
	"server/pkg/errors/validate"
)

func (s *statService) GetByCampaign(ctx context.Context, campaignID string) (*model.Stat, error) {
	if err := validate.IsValidUUID(campaignID); err != nil {
		return nil, validate.NewValidationError("invalid campaign id")
	}
	st, err := s.statRepo.GetByCampaign(ctx, campaignID)
	return st, err
}

func (s *statService) GetByCampaignDaily(ctx context.Context, campaignID string) (*model.Stat, error) {
	if err := validate.IsValidUUID(campaignID); err != nil {
		return nil, validate.NewValidationError("invalid campaign id")
	}
	return s.statRepo.GetByCampaignDaily(ctx, campaignID)
}

func (s *statService) GetByAdvertiser(ctx context.Context, advertiserID string) (*model.Stat, error) {
	if err := validate.IsValidUUID(advertiserID); err != nil {
		return nil, validate.NewValidationError("invalid advertiser id")
	}
	return s.statRepo.GetByAdvertiser(ctx, advertiserID)
}

func (s *statService) GetByAdvertiserDaily(ctx context.Context, advertiserID string) (*model.Stat, error) {
	if err := validate.IsValidUUID(advertiserID); err != nil {
		return nil, validate.NewValidationError("invalid advertiser id")
	}
	return s.statRepo.GetByAdvertiserDaily(ctx, advertiserID)
}
