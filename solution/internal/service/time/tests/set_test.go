package tests

import (
	"context"
	"server/internal/repository"
	"server/internal/repository/mocks"
	"server/internal/service/time"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	t.Parallel()
	type timeRepositoryMockFunc func(mc *minimock.Controller) repository.TimeRepository

	type args struct {
		ctx context.Context
		day int
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		day = gofakeit.Number(5, 100)
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name           string
		args           args
		timeRepository timeRepositoryMockFunc
		wantErr        bool
		err            error
	}{
		{
			name: "Success",
			args: args{
				ctx: ctx,
				day: day,
			},
			timeRepository: func(mc *minimock.Controller) repository.TimeRepository {
				repo := mocks.NewTimeRepositoryMock(mc)
				repo.SetMock.Expect(ctx, day).Return(nil)
				repo.GetMock.Expect(ctx).Return(day, nil)
				repo.UpdateCampaignsStateMock.Expect(ctx).Return(nil)
				return repo
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := tt.timeRepository(mc)
			service := time.NewTimeService(repository)
			err := service.Set(tt.args.ctx, tt.args.day)
			if tt.wantErr {
				require.ErrorIs(t, err, tt.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
