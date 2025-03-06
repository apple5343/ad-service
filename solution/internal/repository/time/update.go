package time

import (
	"context"
	"server/pkg/time"
)

func (r *timeRepository) UpdateCampaignsState(ctx context.Context) error {
	query := `UPDATE campaigns
				SET active = CASE
					WHEN start_date <= ? AND end_date >= ? THEN true
					ELSE false
				END
				WHERE active IS DISTINCT FROM (
					CASE
						WHEN start_date <= ? AND end_date >= ? THEN true
						ELSE false
					END
				);`
	day := time.Day()
	return r.db.WithContext(ctx).Exec(query, day, day, day, day).Error
}
