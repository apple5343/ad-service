package tests

import (
	"context"
	"server/internal/model"
	"server/internal/repository"
	"server/internal/repository/mocks"
	"server/internal/service"
	"server/internal/service/advertiser"
	"testing"

	repo "server/internal/repository/advertiser"
	serviceMocks "server/internal/service/mocks"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	t.Parallel()

	type AdvertiserRepsitoryMock func(mc *minimock.Controller) repository.AdvertiserRepository
	type AiServiceMockFunc func(mc *minimock.Controller) service.AiService

	type args struct {
		ctx     context.Context
		reqBody []*model.Advertiser
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		reqBody = []*model.Advertiser{
			{
				ID:   gofakeit.UUID(),
				Name: gofakeit.Name(),
			},
			{
				ID:   gofakeit.UUID(),
				Name: gofakeit.Name(),
			},
			{
				ID:   gofakeit.UUID(),
				Name: gofakeit.Name(),
			},
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name              string
		args              args
		advertiserRepo    AdvertiserRepsitoryMock
		aiServiceMockFunc AiServiceMockFunc
		err               error
	}{
		{
			name: "Success",
			args: args{
				ctx:     ctx,
				reqBody: reqBody,
			},
			advertiserRepo: func(mc *minimock.Controller) repository.AdvertiserRepository {
				mock := mocks.NewAdvertiserRepositoryMock(mc)
				mock.SaveMock.Expect(ctx, reqBody).Return(nil)
				return mock
			},
			aiServiceMockFunc: func(mc *minimock.Controller) service.AiService {
				return serviceMocks.NewAiServiceMock(mc)
			},
			err: nil,
		},
		{
			"Error conflict",
			args{
				ctx:     ctx,
				reqBody: reqBody,
			},
			func(mc *minimock.Controller) repository.AdvertiserRepository {
				mock := mocks.NewAdvertiserRepositoryMock(mc)
				mock.SaveMock.Expect(ctx, reqBody).Return(repo.ErrAdvertiserAlreadyExists)
				return mock
			},
			func(mc *minimock.Controller) service.AiService {
				return serviceMocks.NewAiServiceMock(mc)
			},
			repo.ErrAdvertiserAlreadyExists,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := tt.advertiserRepo(mc)
			aiService := tt.aiServiceMockFunc(mc)
			service := advertiser.NewAdvertiserService(repository, aiService)

			err := service.Save(tt.args.ctx, tt.args.reqBody)
			if err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}

			require.NoError(t, err)
		})
	}
}
