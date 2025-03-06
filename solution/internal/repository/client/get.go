package client

import (
	"context"
	"server/internal/model"
	"server/internal/repository/client/converter"
	repo "server/internal/repository/client/model"

	"gorm.io/gorm"
)

func (r *clientRepository) Get(ctx context.Context, id string) (*model.Client, error) {
	client := &repo.Client{}
	result := r.db.WithContext(ctx).First(client, "client_id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrClientNotFound
		}
		return nil, result.Error
	}
	return converter.ToClientFromRepo(client), nil
}
