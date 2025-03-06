package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/internal/model"
	repo "server/internal/repository/ads"
	"server/internal/service"
	"server/internal/service/mocks"
	"server/internal/transport/http/handlers/ads"
	"server/internal/transport/http/middleware"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type adsServiceMockFunc func(mc *minimock.Controller) service.AdsService

	type args struct {
		ctx    context.Context
		userID string
	}

	type respBody struct {
		AdID         string `json:"ad_id"`
		AdTitle      string `json:"ad_title"`
		AdText       string `json:"ad_text"`
		AdvertiserID string `json:"advertiser_id"`
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		userID       = gofakeit.UUID()
		Title        = gofakeit.Name()
		Text         = gofakeit.Name()
		AdID         = gofakeit.UUID()
		AdvertiserID = gofakeit.UUID()

		resp = respBody{
			AdID:         AdID,
			AdTitle:      Title,
			AdText:       Text,
			AdvertiserID: AdvertiserID,
		}

		serviceResp = model.Campaign{
			ID:           AdID,
			AdTitle:      Title,
			AdText:       Text,
			AdvertiserID: AdvertiserID,
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name           string
		args           args
		want           respBody
		statusCode     int
		adsServiceMock adsServiceMockFunc
	}{
		{
			name: "Success",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			want:       resp,
			statusCode: 200,
			adsServiceMock: func(mc *minimock.Controller) service.AdsService {
				mock := mocks.NewAdsServiceMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(&serviceResp, nil)
				return mock
			},
		},
		{
			name: "Error not found",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			want:       resp,
			statusCode: 404,
			adsServiceMock: func(mc *minimock.Controller) service.AdsService {
				mock := mocks.NewAdsServiceMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(nil, repo.ErrAdNotFound)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			service := tt.adsServiceMock(mc)
			handler := ads.NewHandler(service)

			req := httptest.NewRequest(http.MethodGet, "/ads?client_id="+tt.args.userID, nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.QueryParams().Add("client_id", tt.args.userID)
			middleware.ErrorHandlerMiddleware(handler.Get())(c)

			require.Equal(t, tt.statusCode, rec.Code)
			if tt.statusCode == http.StatusOK {
				var res respBody
				require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res))
				require.Equal(t, tt.want, res)
			}
		})
	}
}
