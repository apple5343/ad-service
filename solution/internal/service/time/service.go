package time

import (
	"context"
	"server/internal/repository"
	"server/internal/service"
	"server/pkg/time"
)

type timeService struct {
	timeRepository repository.TimeRepository
}

func NewTimeService(timeRepository repository.TimeRepository) service.TimeService {
	s := &timeService{
		timeRepository: timeRepository,
	}
	s.init()
	return s
}

func (s *timeService) init() error {
	day, err := s.timeRepository.Get(context.Background())
	if err != nil {
		return err
	}
	time.Set(day)
	return nil
}
