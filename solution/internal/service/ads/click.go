package ads

import (
	"context"
	"server/pkg/errors/validate"
	metric "server/pkg/metrics"
)

func (s *adsService) Click(ctx context.Context, campaignID, clientID string) error {
	if err := validate.IsValidUUID(campaignID); err != nil {
		return err
	}
	if err := validate.IsValidUUID(clientID); err != nil {
		return err
	}
	ok, err := s.adsRepsitory.IsShownToClient(ctx, campaignID, clientID)
	if err != nil {
		return err
	}
	if !ok {
		return ErrNotShown
	}
	metric.IncAdActionCounter("click")
	return s.adsRepsitory.Click(ctx, campaignID, clientID)
}
