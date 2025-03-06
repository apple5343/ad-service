package advertiser

import (
	"context"
	"server/internal/model"
	"server/internal/repository/advertiser/converter"
	"server/pkg/client/postgres"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *advertiserRepository) AddScore(ctx context.Context, score *model.Score) error {
	res := r.db.WithContext(ctx).Save(converter.ToRepoFromScore(score))
	if res.Error != nil {
		if pqErr, ok := res.Error.(*pgconn.PgError); ok {
			if pqErr.Code == postgres.CodeForeignKeyViolation {
				return ErrAdvertiserOrClientNotFound
			}
		}
		return res.Error
	}
	return nil
}
