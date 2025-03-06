package tests

import (
	"context"
	"server/internal/model"
	"server/internal/repository"
	"server/internal/repository/mocks"
	"server/internal/service/campaign"
	"testing"

	repo "server/internal/repository/campaign"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type CampaignRepositoryMock func(mc *minimock.Controller) repository.CampaignRepository

	type args struct {
		ctx          context.Context
		campaignID   string
		advertiserID string
	}

	var (
		ctx          = context.Background()
		mc           = minimock.NewController(t)
		campaignID   = gofakeit.UUID()
		advertiserID = gofakeit.UUID()

		serviceResonse = &model.Campaign{
			ID:                campaignID,
			AdvertiserID:      advertiserID,
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
		name               string
		args               args
		campaignRepository CampaignRepositoryMock
		want               *model.Campaign
		err                error
	}{
		{
			name: "Success",
			args: args{
				ctx:          ctx,
				campaignID:   campaignID,
				advertiserID: advertiserID,
			},
			campaignRepository: func(mc *minimock.Controller) repository.CampaignRepository {
				mock := mocks.NewCampaignRepositoryMock(mc)
				mock.GetMock.Expect(ctx, advertiserID, campaignID).Return(serviceResonse, nil)
				return mock
			},
			want: serviceResonse,
			err:  nil,
		},
		{
			name: "Error campaign not found",
			args: args{
				ctx:          ctx,
				campaignID:   campaignID,
				advertiserID: advertiserID,
			},
			campaignRepository: func(mc *minimock.Controller) repository.CampaignRepository {
				mock := mocks.NewCampaignRepositoryMock(mc)
				mock.GetMock.Expect(ctx, advertiserID, campaignID).Return(nil, repo.ErrCampaignNotFound)
				return mock
			},
			want: nil,
			err:  repo.ErrCampaignNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := tt.campaignRepository(mc)
			service := campaign.NewCampaignService(repository, nil, nil)
			campaign, err := service.Get(tt.args.ctx, tt.args.advertiserID, tt.args.campaignID)
			if err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, campaign)
		})
	}
}

func TestList(t *testing.T) {
	t.Parallel()

	type CampaignRepositoryMock func(mc *minimock.Controller) repository.CampaignRepository

	type args struct {
		ctx          context.Context
		advertiserID string
	}

	var (
		ctx            = context.Background()
		mc             = minimock.NewController(t)
		advertiserID   = gofakeit.UUID()
		page           = gofakeit.Number(1, 10)
		size           = 2
		serviceResonse = []*model.Campaign{
			{
				ID:                gofakeit.UUID(),
				AdvertiserID:      advertiserID,
				AdTitle:           gofakeit.Name(),
				AdText:            gofakeit.Name(),
				ClicksLimit:       gofakeit.Number(0, 30),
				ImpressionsLimit:  gofakeit.Number(40, 100),
				CostPerClick:      float64(gofakeit.Number(0, 100)),
				CostPerImpression: float64(gofakeit.Number(0, 100)),
				StartDate:         gofakeit.Number(0, 20),
				EndDate:           gofakeit.Number(40, 60),
			},
			{
				ID:                gofakeit.UUID(),
				AdvertiserID:      advertiserID,
				AdTitle:           gofakeit.Name(),
				AdText:            gofakeit.Name(),
				ClicksLimit:       gofakeit.Number(0, 30),
				ImpressionsLimit:  gofakeit.Number(40, 100),
				CostPerClick:      float64(gofakeit.Number(0, 100)),
				CostPerImpression: float64(gofakeit.Number(0, 100)),
				StartDate:         gofakeit.Number(0, 20),
				EndDate:           gofakeit.Number(40, 60),
			},
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		campaignRepository CampaignRepositoryMock
		want               []*model.Campaign
		err                error
	}{
		{
			name: "Success",
			args: args{
				ctx:          ctx,
				advertiserID: advertiserID,
			},
			campaignRepository: func(mc *minimock.Controller) repository.CampaignRepository {
				mock := mocks.NewCampaignRepositoryMock(mc)
				mock.ListMock.Expect(ctx, advertiserID, page, size).Return(serviceResonse, nil)
				return mock
			},
			want: serviceResonse,
			err:  nil,
		},
		{
			name: "Error advertiser not found",
			args: args{
				ctx:          ctx,
				advertiserID: advertiserID,
			},
			campaignRepository: func(mc *minimock.Controller) repository.CampaignRepository {
				mock := mocks.NewCampaignRepositoryMock(mc)
				mock.ListMock.Expect(ctx, advertiserID, page, size).Return(nil, repo.ErrAdvertiserNotRegistered)
				return mock
			},
			want: nil,
			err:  repo.ErrAdvertiserNotRegistered,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := tt.campaignRepository(mc)
			service := campaign.NewCampaignService(repository, nil, nil)
			campaigns, err := service.List(tt.args.ctx, tt.args.advertiserID, page, size)
			if err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.want, campaigns)
		})
	}
}
