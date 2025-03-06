package campaign

import (
	"context"
	"fmt"
	"server/internal/model"
	"server/pkg/errors"
)

func (s *campaignService) SetModerateStatus(_ context.Context, moderate bool) error {
	s.moderate = moderate
	return nil
}

func (s *campaignService) Moderate(ctx context.Context, campaign *model.Campaign) error {
	ok, attributes, err := s.ai.ModerateText(ctx, campaign.AdText)
	if err != nil {
		return err
	}
	if !ok {
		return errors.NewError(fmt.Sprintf("ad text is not approved: %v", attributes), errors.BadRequest)
	}

	ok, attributes, err = s.ai.ModerateText(ctx, campaign.AdTitle)
	if err != nil {
		return err
	}
	if !ok {
		return errors.NewError(fmt.Sprintf("ad title is not approved: %v", attributes), errors.BadRequest)
	}
	return nil
}
