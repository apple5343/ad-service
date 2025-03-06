package advertiser

import (
	"context"
	"server/internal/model"
	"server/pkg/errors/validate"
)

func (s *advertiserService) AddScore(ctx context.Context, score *model.Score) error {
	if err := validate.IsValidUUID(score.AdvertiserID); err != nil {
		return validate.NewValidationError("invalid advertiser id")
	}
	if err := validate.IsValidUUID(score.ClientID); err != nil {
		return validate.NewValidationError("invalid client id")
	}
	if score != nil && score.Score < 0 {
		return ErrInvalidScore
	}
	return s.repo.AddScore(ctx, score)
}
