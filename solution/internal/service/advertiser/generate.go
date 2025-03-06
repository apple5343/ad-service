package advertiser

import (
	"context"
	"server/pkg/errors/validate"
)

func (s *advertiserService) Generate(ctx context.Context, adveristerID, campaignName string) (string, error) {
	if err := validate.IsValidUUID(adveristerID); err != nil {
		return "", err
	}

	adverister, err := s.Get(ctx, adveristerID)
	if err != nil {
		return "", err
	}

	description, err := s.aiService.GenerateCampaignDescription(ctx, adverister.Name, campaignName)
	if err != nil {
		return "", err
	}
	return description, nil

}
