package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	repo "server/internal/repository/ads"
	"server/internal/service"
	"server/internal/service/mocks"
	httpapi "server/internal/transport/http"
	"server/internal/transport/http/handlers/ads"
	"server/internal/transport/http/middleware"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestClick(t *testing.T) {
	t.Parallel()
	type adsServiceMockFunc func(mc *minimock.Controller) service.AdsService

	type reqBody struct {
		ClientID string `json:"client_id"`
	}

	type args struct {
		ctx  context.Context
		req  reqBody
		adID string
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		clientID = gofakeit.UUID()
		adID     = gofakeit.UUID()

		req = reqBody{
			ClientID: clientID,
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name           string
		args           args
		statusCode     int
		adsServiceMock adsServiceMockFunc
	}{
		{
			name: "Success",
			args: args{
				ctx:  ctx,
				req:  req,
				adID: adID,
			},
			statusCode: 204,
			adsServiceMock: func(mc *minimock.Controller) service.AdsService {
				mock := mocks.NewAdsServiceMock(mc)
				mock.ClickMock.Expect(ctx, adID, clientID).Return(nil)
				return mock
			},
		},
		{
			name: "Error not found",
			args: args{
				ctx:  ctx,
				req:  req,
				adID: adID,
			},
			statusCode: 404,
			adsServiceMock: func(mc *minimock.Controller) service.AdsService {
				mock := mocks.NewAdsServiceMock(mc)
				mock.ClickMock.Expect(ctx, adID, clientID).Return(repo.ErrAdNotFound)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			e.Validator = &httpapi.CustomValidator{}

			service := tt.adsServiceMock(mc)
			handler := ads.NewHandler(service)
			reqB, _ := json.Marshal(tt.args.req)
			req := httptest.NewRequest(http.MethodGet, "/ads/"+tt.args.adID+"/click", bytes.NewBuffer(reqB))
			req.Header.Set("Content-Type", "application/json")

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("adId")
			c.SetParamValues(tt.args.adID)

			middleware.ErrorHandlerMiddleware(handler.Click())(c)

			require.Equal(t, tt.statusCode, rec.Code)
		})
	}
}
