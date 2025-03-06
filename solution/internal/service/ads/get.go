package ads

import (
	"context"
	"server/internal/model"
	"server/pkg/errors/validate"
	metric "server/pkg/metrics"
)

func (s *adsService) Get(ctx context.Context, clientID string) (*model.Campaign, error) {
	if err := validate.IsValidUUID(clientID); err != nil {
		return nil, err
	}
	campaignID, _, cost, err := s.adsRepsitory.Get(ctx, clientID)
	if err != nil {
		return nil, err
	}
	campaign, err := s.adsRepsitory.Impression(ctx, campaignID, clientID, cost)
	if err != nil {
		return nil, err
	}
	metric.IncAdActionCounter("impression")
	return campaign, nil
}
