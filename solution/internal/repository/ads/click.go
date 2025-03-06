package ads

import (
	"context"
	"server/pkg/client/postgres"
	"server/pkg/time"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func (r *adsRepsitory) Click(ctx context.Context, campaignID, clientID string) error {
	type Click struct {
		CampaignID string
		ClientID   string
		Cost       float64
		Date       int
	}

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var cost float64
		query := `UPDATE campaigns SET click_count = click_count + 1 WHERE campaign_id = ? AND active = true RETURNING cost_per_click`
		res := tx.Raw(query, campaignID).Scan(&cost)

		if res.Error != nil {
			return res.Error
		}

		if res.RowsAffected == 0 {
			return ErrEndedAd
		}

		click := Click{
			CampaignID: campaignID,
			ClientID:   clientID,
			Cost:       cost,
			Date:       time.Day(),
		}

		if err := tx.Create(&click).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrCampaignNotFound
		}
		if pqErr, ok := err.(*pgconn.PgError); ok {
			if pqErr.Code == postgres.CodeUniqueViolation {
				return nil
			}
		}
		return err
	}
	return nil
}
