package tests

import (
	"context"
	"server/internal/repository"
	"server/internal/repository/mocks"
	"server/internal/service/campaign"
	"testing"

	repo "server/internal/repository/campaign"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
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
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		campaignRepository CampaignRepositoryMock
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
				mock.DeleteMock.Expect(ctx, advertiserID, campaignID).Return(nil)
				return mock
			},
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
				mock.DeleteMock.Expect(ctx, advertiserID, campaignID).Return(repo.ErrCampaignNotFound)
				return mock
			},
			err: repo.ErrCampaignNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := tt.campaignRepository(mc)
			service := campaign.NewCampaignService(repository, nil, nil)
			err := service.Delete(tt.args.ctx, tt.args.advertiserID, tt.args.campaignID)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
