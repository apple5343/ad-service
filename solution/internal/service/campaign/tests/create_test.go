package tests

import (
	"context"
	"server/internal/model"
	"server/internal/repository"
	"server/internal/repository/mocks"
	"server/internal/service"
	"server/internal/service/campaign"
	"testing"

	repo "server/internal/repository/campaign"
	serviceMock "server/internal/service/mocks"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type CampaignRepositoryMock func(mc *minimock.Controller) repository.CampaignRepository
	type CampaignObjectRepositoryMock func(mc *minimock.Controller) repository.CampaignObjectRepository

	type AiServiceMockFunc func(mc *minimock.Controller) service.AiService

	type args struct {
		ctx     context.Context
		reqBody *model.Campaign
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		reqBody = model.Campaign{
			ID:                gofakeit.UUID(),
			AdvertiserID:      gofakeit.UUID(),
			AdTitle:           gofakeit.Name(),
			AdText:            gofakeit.Name(),
			ClicksLimit:       gofakeit.Number(0, 30),
			ImpressionsLimit:  gofakeit.Number(40, 100),
			CostPerClick:      float64(gofakeit.Number(0, 100)),
			CostPerImpression: float64(gofakeit.Number(0, 100)),
			StartDate:         gofakeit.Number(0, 20),
			EndDate:           gofakeit.Number(40, 60),
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name                         string
		args                         args
		want                         *model.Campaign
		campaignRepository           CampaignRepositoryMock
		aiService                    AiServiceMockFunc
		campaignObjectRepositoryMock CampaignObjectRepositoryMock
		err                          error
	}{
		{
			name: "Success",
			args: args{
				ctx:     ctx,
				reqBody: &reqBody,
			},
			campaignRepository: func(mc *minimock.Controller) repository.CampaignRepository {
				mock := mocks.NewCampaignRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, &reqBody).Return(&reqBody, nil)
				return mock
			},
			aiService: func(mc *minimock.Controller) service.AiService {
				return serviceMock.NewAiServiceMock(mc)
			},
			campaignObjectRepositoryMock: func(mc *minimock.Controller) repository.CampaignObjectRepository {
				return mocks.NewCampaignObjectRepositoryMock(mc)
			},
			want: &reqBody,
		},
		{
			name: "Error advertiser not registered",
			args: args{
				ctx:     ctx,
				reqBody: &reqBody,
			},
			campaignRepository: func(mc *minimock.Controller) repository.CampaignRepository {
				mock := mocks.NewCampaignRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, &reqBody).Return(nil, repo.ErrAdvertiserNotRegistered)
				return mock
			},
			aiService: func(mc *minimock.Controller) service.AiService {
				return serviceMock.NewAiServiceMock(mc)
			},
			campaignObjectRepositoryMock: func(mc *minimock.Controller) repository.CampaignObjectRepository {
				return mocks.NewCampaignObjectRepositoryMock(mc)
			},
			err: repo.ErrAdvertiserNotRegistered,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := tt.campaignRepository(mc)
			objectRepository := tt.campaignObjectRepositoryMock(mc)
			aiService := tt.aiService(mc)
			service := campaign.NewCampaignService(repository, objectRepository, aiService)

			campaign, err := service.Create(tt.args.ctx, tt.args.reqBody)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, campaign)
		})
	}

}
