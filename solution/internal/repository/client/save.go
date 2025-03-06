package client

import (
	"context"
	"server/internal/model"
	"server/internal/repository/client/converter"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *clientRepository) Save(ctx context.Context, clients []*model.Client) error {
	result := r.db.WithContext(ctx).Save(converter.ToReposFromClients(clients))
	if result.Error != nil {
		if pqErr, ok := result.Error.(*pgconn.PgError); ok {
			if pqErr.Code == "21000" {
				return ErrClientAlreadyExists
			}
		}
		return result.Error
	}
	return nil
}
