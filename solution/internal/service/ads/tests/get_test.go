package tests

import (
	"context"
	"server/internal/model"
	"server/internal/repository"
	"server/internal/repository/mocks"
	"server/internal/service/ads"
	"testing"

	repo "server/internal/repository/ads"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type AdsRepsitoryMock func(mc *minimock.Controller) repository.AdsRepsitory

	type args struct {
		ctx      context.Context
		clientID string
	}

	var (
		ctx            = context.Background()
		mc             = minimock.NewController(t)
		campaignID     = gofakeit.UUID()
		serviceResonse = model.Campaign{
			ID:           campaignID,
			AdvertiserID: gofakeit.UUID(),
			AdTitle:      gofakeit.Name(),
			AdText:       gofakeit.Name(),
		}
		clientID = gofakeit.UUID()

		score = gofakeit.Number(0, 100)
		cost  = float64(gofakeit.Number(0, 100))
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name          string
		args          args
		adsRepository AdsRepsitoryMock
		want          model.Campaign
		err           error
	}{
		{
			name: "Success",
			args: args{
				ctx:      ctx,
				clientID: clientID,
			},
			adsRepository: func(mc *minimock.Controller) repository.AdsRepsitory {
				mock := mocks.NewAdsRepsitoryMock(mc)
				mock.GetMock.Expect(ctx, clientID).Return(campaignID, score, cost, nil)
				mock.ImpressionMock.Expect(ctx, campaignID, clientID, cost).Return(&serviceResonse, nil)
				return mock
			},
			want: serviceResonse,
		},
		{
			name: "Ad not found",
			args: args{
				ctx:      ctx,
				clientID: clientID,
			},
			adsRepository: func(mc *minimock.Controller) repository.AdsRepsitory {
				mock := mocks.NewAdsRepsitoryMock(mc)
				mock.GetMock.Expect(ctx, clientID).Return("", 0, 0, repo.ErrAdNotFound)
				return mock
			},
			err: repo.ErrAdNotFound,
		},
		{
			name: "Client not registered",
			args: args{
				ctx:      ctx,
				clientID: clientID,
			},
			adsRepository: func(mc *minimock.Controller) repository.AdsRepsitory {
				mock := mocks.NewAdsRepsitoryMock(mc)
				mock.GetMock.Expect(ctx, clientID).Return("", 0, 0, repo.ErrClientNotRegistered)
				return mock
			},
			err: repo.ErrClientNotRegistered,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := tt.adsRepository(mc)
			service := ads.NewAdsService(repository)
			campaign, err := service.Get(tt.args.ctx, tt.args.clientID)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, *campaign)
		})
	}
}
