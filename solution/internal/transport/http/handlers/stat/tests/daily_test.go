package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/internal/model"
	"server/internal/repository/advertiser"
	"server/internal/repository/campaign"
	"server/internal/service"
	"server/internal/service/mocks"
	"server/internal/transport/http/handlers/stat"
	"server/internal/transport/http/middleware"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestGetByCmampaignDaily(t *testing.T) {
	t.Parallel()
	type statServiceMockFunc func(mc *minimock.Controller) service.StatService

	type respBody struct {
		ImpressionsCount int     `json:"impressions_count"`
		ClicksCount      int     `json:"clicks_count"`
		Coversion        float64 `json:"conversion"`
		SpentImpressions float64 `json:"spent_impressions"`
		SpentClicks      float64 `json:"spent_clicks"`
		SpentTotal       float64 `json:"spent_total"`
	}

	type args struct {
		ctx        context.Context
		campaignID string
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		campaignID       = gofakeit.UUID()
		impressionsCount = gofakeit.Number(1, 100)
		clicksCount      = gofakeit.Number(1, 100)
		spentImpressions = float64(gofakeit.Number(0, 100))
		spentClicks      = float64(gofakeit.Number(0, 100))
		conversion       = float64(clicksCount) / float64(impressionsCount) * 100
		spentTotal       = spentImpressions + spentClicks

		resp = respBody{
			ImpressionsCount: impressionsCount,
			ClicksCount:      clicksCount,
			Coversion:        conversion,
			SpentImpressions: spentImpressions,
			SpentClicks:      spentClicks,
			SpentTotal:       spentTotal,
		}

		serviceResp = model.Stat{
			ImpressionsCount: impressionsCount,
			ClicksCount:      clicksCount,
			Conversion:       conversion,
			SpentImpressions: spentImpressions,
			SpentClicks:      spentClicks,
			SpentTotal:       spentTotal,
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            respBody
		statusCode      int
		statServiceMock statServiceMockFunc
	}{
		{
			name: "Success",
			args: args{
				ctx:        ctx,
				campaignID: campaignID,
			},
			want:       resp,
			statusCode: 200,
			statServiceMock: func(mc *minimock.Controller) service.StatService {
				mock := mocks.NewStatServiceMock(mc)
				mock.GetByCampaignDailyMock.Expect(ctx, campaignID).Return(&serviceResp, nil)
				return mock
			},
		},
		{
			name: "Err not found",
			args: args{
				ctx:        ctx,
				campaignID: campaignID,
			},
			want:       resp,
			statusCode: 400,
			statServiceMock: func(mc *minimock.Controller) service.StatService {
				mock := mocks.NewStatServiceMock(mc)
				mock.GetByCampaignDailyMock.Expect(ctx, campaignID).Return(nil, campaign.ErrCampaignNotFound)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			service := tt.statServiceMock(mc)
			handler := stat.NewHandler(service)

			req := httptest.NewRequest(http.MethodGet, "/stat/campaigns/"+tt.args.campaignID+"/daily", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetParamNames("campaignId")
			c.SetParamValues(tt.args.campaignID)

			middleware.ErrorHandlerMiddleware(handler.GetByCampaignDaily())(c)

			require.Equal(t, tt.statusCode, rec.Code)
			if tt.statusCode == http.StatusOK {
				var res respBody
				require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res))
				require.Equal(t, tt.want, res)
			}
		})
	}
}

func TestGetByAdvertiserDaily(t *testing.T) {
	t.Parallel()
	type statServiceMockFunc func(mc *minimock.Controller) service.StatService

	type respBody struct {
		ImpressionsCount int     `json:"impressions_count"`
		ClicksCount      int     `json:"clicks_count"`
		Coversion        float64 `json:"conversion"`
		SpentImpressions float64 `json:"spent_impressions"`
		SpentClicks      float64 `json:"spent_clicks"`
		SpentTotal       float64 `json:"spent_total"`
	}

	type args struct {
		ctx          context.Context
		advertiserID string
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		advertiserID     = gofakeit.UUID()
		impressionsCount = gofakeit.Number(1, 100)
		clicksCount      = gofakeit.Number(1, 100)
		spentImpressions = float64(gofakeit.Number(0, 100))
		spentClicks      = float64(gofakeit.Number(0, 100))
		conversion       = float64(clicksCount) / float64(impressionsCount) * 100
		spentTotal       = spentImpressions + spentClicks

		resp = respBody{
			ImpressionsCount: impressionsCount,
			ClicksCount:      clicksCount,
			Coversion:        conversion,
			SpentImpressions: spentImpressions,
			SpentClicks:      spentClicks,
			SpentTotal:       spentTotal,
		}

		serviceResp = model.Stat{
			ImpressionsCount: impressionsCount,
			ClicksCount:      clicksCount,
			Conversion:       conversion,
			SpentImpressions: spentImpressions,
			SpentClicks:      spentClicks,
			SpentTotal:       spentTotal,
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            respBody
		statusCode      int
		statServiceMock statServiceMockFunc
	}{
		{
			name: "Success",
			args: args{
				ctx:          ctx,
				advertiserID: advertiserID,
			},
			want:       resp,
			statusCode: 200,
			statServiceMock: func(mc *minimock.Controller) service.StatService {
				mock := mocks.NewStatServiceMock(mc)
				mock.GetByAdvertiserDailyMock.Expect(ctx, advertiserID).Return(&serviceResp, nil)
				return mock
			},
		},
		{
			name: "Err not found",
			args: args{
				ctx:          ctx,
				advertiserID: advertiserID,
			},
			want:       resp,
			statusCode: 400,
			statServiceMock: func(mc *minimock.Controller) service.StatService {
				mock := mocks.NewStatServiceMock(mc)
				mock.GetByAdvertiserDailyMock.Expect(ctx, advertiserID).Return(nil, advertiser.ErrAdvertiserNotFound)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			service := tt.statServiceMock(mc)
			handler := stat.NewHandler(service)

			req := httptest.NewRequest(http.MethodGet, "/stat/advertiser/"+tt.args.advertiserID+"/daily", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetParamNames("advertiserId")
			c.SetParamValues(tt.args.advertiserID)

			middleware.ErrorHandlerMiddleware(handler.GetByAdvertiserDaily())(c)

			require.Equal(t, tt.statusCode, rec.Code)
			if tt.statusCode == http.StatusOK {
				var res respBody
				require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &res))
				require.Equal(t, tt.want, res)
			}
		})
	}
}
