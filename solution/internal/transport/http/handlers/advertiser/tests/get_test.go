package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"server/internal/model"
	"server/internal/service"
	"server/internal/service/mocks"
	"server/internal/transport/http/handlers/advertiser"
	"server/internal/transport/http/middleware"
	"testing"

	repo "server/internal/repository/advertiser"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type advertiserServiceMockFunc func(mc *minimock.Controller) service.AdvertiserService

	type args struct {
		ctx   context.Context
		advID string
	}

	type respBody struct {
		AdvertiserID string `json:"advertiser_id"`
		Name         string `json:"name"`
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		advID = gofakeit.UUID()
		name  = gofakeit.Name()

		serviceResp = model.Advertiser{
			ID:   advID,
			Name: name,
		}

		resp = respBody{
			AdvertiserID: advID,
			Name:         name,
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name                  string
		args                  args
		advertiserServiceMock advertiserServiceMockFunc
		err                   error
		statusCode            int
		want                  respBody
	}{
		{
			name: "Success",
			args: args{
				ctx:   ctx,
				advID: advID,
			},
			advertiserServiceMock: func(mc *minimock.Controller) service.AdvertiserService {
				mock := mocks.NewAdvertiserServiceMock(mc)
				mock.GetMock.Expect(ctx, advID).Return(&serviceResp, nil)
				return mock
			},
			statusCode: 200,
			want:       resp,
		},
		{
			name: "Error not found",
			args: args{
				ctx:   ctx,
				advID: advID,
			},
			advertiserServiceMock: func(mc *minimock.Controller) service.AdvertiserService {
				mock := mocks.NewAdvertiserServiceMock(mc)
				mock.GetMock.Expect(ctx, advID).Return(nil, repo.ErrAdvertiserNotFound)
				return mock
			},
			statusCode: 400,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			service := tt.advertiserServiceMock(mc)
			handler := advertiser.NewHandler(service)

			req := httptest.NewRequest(http.MethodGet, "/advertisers/:id", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			c.SetParamNames("advertiserId")
			c.SetParamValues(tt.args.advID)
			middleware.ErrorHandlerMiddleware(handler.Get())(c)

			require.Equal(t, tt.statusCode, rec.Code)
			if tt.statusCode == http.StatusOK {
				var res respBody
				require.NoError(t, json.NewDecoder(rec.Body).Decode(&res))
				require.Equal(t, tt.want, res)
			}
		})
	}
}
