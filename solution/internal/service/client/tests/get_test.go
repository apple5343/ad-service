package tests

import (
	"context"
	"server/internal/model"
	"server/internal/repository"
	"server/internal/repository/mocks"
	"server/internal/service/client"
	"testing"

	repo "server/internal/repository/client"

	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type ClientRepositoryMock func(mc *minimock.Controller) repository.ClientRepository

	type args struct {
		ctx      context.Context
		clientID string
	}

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		clientID   = gofakeit.UUID()
		clientResp = &model.Client{
			ID:    clientID,
			Login: gofakeit.Name(),
		}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name             string
		args             args
		clientRepository ClientRepositoryMock
		want             *model.Client
		err              error
	}{
		{
			name: "Success",
			args: args{
				ctx:      ctx,
				clientID: clientID,
			},
			clientRepository: func(mc *minimock.Controller) repository.ClientRepository {
				mock := mocks.NewClientRepositoryMock(mc)
				mock.GetMock.Expect(ctx, clientID).Return(clientResp, nil)
				return mock
			},
			err:  nil,
			want: clientResp,
		},
		{
			name: "Error client not found",
			args: args{
				ctx:      ctx,
				clientID: clientID,
			},
			clientRepository: func(mc *minimock.Controller) repository.ClientRepository {
				mock := mocks.NewClientRepositoryMock(mc)
				mock.GetMock.Expect(ctx, clientID).Return(nil, repo.ErrClientNotFound)
				return mock
			},
			err:  repo.ErrClientNotFound,
			want: nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repository := tt.clientRepository(mc)
			service := client.NewClientService(repository)

			cl, err := service.Get(tt.args.ctx, tt.args.clientID)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.want, cl)
		})
	}
}
