package tests

import (
	"context"
	"server/internal/repository"
	repo "server/internal/repository/ads"
	"server/internal/repository/mocks"
	"server/internal/service/ads"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestClick(t *testing.T) {
	t.Parallel()

	type AdsRepsitoryMock func(mc *minimock.Controller) repository.AdsRepsitory

	type args struct {
		ctx      context.Context
		adID     string
		clientID string
	}
	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		adID     = gofakeit.UUID()
		clientID = gofakeit.UUID()
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name          string
		args          args
		adsRepository AdsRepsitoryMock
		err           error
	}{
		{
			name: "Success",
			args: args{
				ctx:      ctx,
				adID:     adID,
				clientID: clientID,
			},
			adsRepository: func(mc *minimock.Controller) repository.AdsRepsitory {
				mock := mocks.NewAdsRepsitoryMock(mc)
				mock.IsShownToClientMock.Expect(ctx, adID, clientID).Return(true, nil)
				mock.ClickMock.Expect(ctx, adID, clientID).Return(nil)
				return mock
			},
		},
		{
			name: "Error not shown to client",
			args: args{
				ctx:      ctx,
				adID:     adID,
				clientID: clientID,
			},
			adsRepository: func(mc *minimock.Controller) repository.AdsRepsitory {
				mock := mocks.NewAdsRepsitoryMock(mc)
				mock.IsShownToClientMock.Expect(ctx, adID, clientID).Return(false, nil)
				return mock
			},
			err: ads.ErrNotShown,
		},
		{
			name: "Error client not registered",
			args: args{
				ctx:      ctx,
				adID:     adID,
				clientID: clientID,
			},
			adsRepository: func(mc *minimock.Controller) repository.AdsRepsitory {
				mock := mocks.NewAdsRepsitoryMock(mc)
				mock.IsShownToClientMock.Expect(ctx, adID, clientID).Return(false, repo.ErrClientNotRegistered)
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
			err := service.Click(tt.args.ctx, tt.args.adID, tt.args.clientID)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
