package tests

import (
	"bytes"
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

	httpapi "server/internal/transport/http"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	t.Parallel()
	type advertiserServiceMockFunc func(mc *minimock.Controller) service.AdvertiserService

	type body struct {
		AdvertiserID string `json:"advertiser_id"`
		Name         string `json:"name"`
	}
	type args struct {
		ctx     context.Context
		reqBody []*body
	}

	serviceResp := []*model.Advertiser{
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

	var req []*body
	for _, v := range serviceResp {
		req = append(req, &body{
			AdvertiserID: v.ID,
			Name:         v.Name,
		})
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name                  string
		args                  args
		advertiserServiceMock advertiserServiceMockFunc
		want                  []*body
		statusCode            int
		err                   error
	}{
		{
			name: "Success",
			args: args{
				ctx:     ctx,
				reqBody: req,
			},
			advertiserServiceMock: func(mc *minimock.Controller) service.AdvertiserService {
				mock := mocks.NewAdvertiserServiceMock(mc)
				mock.SaveMock.Return(nil)
				return mock
			},
			want:       req,
			statusCode: http.StatusCreated,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			e := echo.New()
			e.Validator = &httpapi.CustomValidator{}

			service := tt.advertiserServiceMock(mc)
			handler := advertiser.NewHandler(service)

			reqB, _ := json.Marshal(tt.args.reqBody)
			req := httptest.NewRequest(http.MethodPost, "/advertisers/bulk", bytes.NewReader(reqB))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			middleware.ErrorHandlerMiddleware(handler.Save())(c)

			require.Equal(t, tt.statusCode, rec.Code)
			if tt.statusCode == http.StatusOK {
				var got []*body
				err := json.Unmarshal(rec.Body.Bytes(), &got)
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			}
		})
	}
}
