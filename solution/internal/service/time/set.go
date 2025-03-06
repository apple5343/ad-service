package time

import (
	"context"
	"server/pkg/logger"
	"server/pkg/time"

	"go.uber.org/zap"
)

func (s *timeService) Set(ctx context.Context, day int) error {
	if day < 0 {
		return ErrInvalidDay
	}
	if err := s.timeRepository.Set(ctx, day); err != nil {
		return err
	}
	time.Set(day)
	logger.Info("Time was set", zap.Int("day", day))
	return s.timeRepository.UpdateCampaignsState(ctx)
}
