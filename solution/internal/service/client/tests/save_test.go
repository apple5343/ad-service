package tests

import (
	"context"
	"server/internal/model"
	"server/internal/repository"
	"server/internal/repository/mocks"
	"server/internal/service/client"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestSave(t *testing.T) {
	t.Parallel()

	type ClientRepositoryMock func(mc *minimock.Controller) repository.ClientRepository

	type args struct {
		ctx  context.Context
		body []*model.Client
	}
	genders := []string{"MALE", "FEMALE"}
	body := []*model.Client{}
	for i := 0; i < 10; i++ {
		body = append(body, &model.Client{
			ID:       gofakeit.UUID(),
			Login:    gofakeit.Name(),
			Age:      gofakeit.Number(0, 100),
			Location: gofakeit.City(),
			Gender:   genders[gofakeit.Number(0, 1)],
		})
	}

	invalidBody := []*model.Client{}
	for i := 0; i < 10; i++ {
		invalidBody = append(invalidBody, &model.Client{
			Login:    gofakeit.Name(),
			Age:      gofakeit.Number(0, 100),
			Location: gofakeit.City(),
			Gender:   genders[gofakeit.Number(0, 1)],
		})
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name             string
		args             args
		clientRepository ClientRepositoryMock
		wantErr          bool
	}{
		{
			name: "Success",
			args: args{
				ctx:  ctx,
				body: body,
			},
			clientRepository: func(mc *minimock.Controller) repository.ClientRepository {
				mock := mocks.NewClientRepositoryMock(mc)
				mock.SaveMock.Expect(ctx, body).Return(nil)
				return mock
			},
		},
		{
			name: "Error invalid uuid",
			args: args{
				ctx:  ctx,
				body: invalidBody,
			},
			clientRepository: func(mc *minimock.Controller) repository.ClientRepository {
				mock := mocks.NewClientRepositoryMock(mc)
				return mock
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := tt.clientRepository(mc)
			service := client.NewClientService(repository)

			err := service.Save(tt.args.ctx, tt.args.body)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
