package ads

import (
	"context"
	"server/internal/model"
	"server/internal/repository/ads/converter"
	repo "server/internal/repository/campaign/model"
	repoClient "server/internal/repository/client/model"
	"server/pkg/client/postgres"
	"server/pkg/time"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

const (
	a = 0.1
	b = 0.5
)

func (r *adsRepsitory) Get(ctx context.Context, clientID string) (string, int, float64, error) {
	client := repoClient.Client{}
	res := r.db.WithContext(ctx).First(&client, "client_id = ?", clientID)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return "", 0, 0, ErrClientNotRegistered
		}
		return "", 0, 0, res.Error
	}

	query := `SELECT campaign_id, cost_per_impression as cost,
				(? * COALESCE(s.score, 0) + ? * (cost_per_click + cost_per_impression)) * 
				(
					CASE
						WHEN c.clicks_limit = 0 THEN 1
						ELSE (1 - c.click_count / CAST(c.clicks_limit AS FLOAT))
					END + 
					CASE
						WHEN c.impressions_limit = 0 THEN 1
						ELSE (1 - c.impression_count / CAST(c.impressions_limit AS FLOAT))
					END
				) / 2 AS total_score, s.score
				FROM campaigns c
				LEFT JOIN scores s ON c.advertiser_id = s.advertiser_id AND s.client_id = ?
				WHERE c.active = true AND CASE WHEN c.impressions_limit = 0 THEN 2 ELSE c.impression_count / CAST(c.impressions_limit AS FLOAT) END < 1.03
				AND (c.age_from <= ? OR c.age_from IS NULL) AND (c.age_to >= ? OR c.age_to IS NULL)
				AND (c.gender = ? OR c.gender = 'ALL' OR c.gender IS NULL)
				AND (? = c.location OR c.location IS NULL)
				ORDER BY total_score DESC
				LIMIT 1;`

	var result struct {
		CampaignID string
		Score      int
		TotalScore float32
		Cost       float64
	}
	res = r.db.WithContext(ctx).Raw(query, a, b, clientID, client.Age, client.Age, client.Gender, client.Location).Scan(&result)
	if res.Error != nil {
		return "", 0, 0, res.Error
	}
	if res.RowsAffected == 0 {
		return "", 0, 0, ErrAdNotFound
	}
	return result.CampaignID, result.Score, result.Cost, nil
}

func (r *adsRepsitory) Impression(ctx context.Context, campaignID string, clientID string, cost float64) (*model.Campaign, error) {
	type Impression struct {
		CampaignID string
		ClientID   string
		Cost       float64
		Date       int
	}

	impression := Impression{
		CampaignID: campaignID,
		ClientID:   clientID,
		Cost:       cost,
		Date:       time.Day(),
	}

	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&impression).Error; err != nil {
			return err
		}

		if err := tx.Model(&repo.Campaign{}).
			Where("campaign_id = ?", campaignID).
			Update("impression_count", gorm.Expr("impression_count + 1")).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			if pqErr.Code != postgres.CodeUniqueViolation {
				return nil, err
			}
		}
	}
	campaign := &repo.Campaign{}
	if err := r.db.WithContext(ctx).Where("campaign_id = ?", campaignID).First(campaign).Error; err != nil {
		return nil, err
	}
	return converter.ToCampaignFromRepo(campaign), nil
}

func (r *adsRepsitory) IsShownToClient(ctx context.Context, campaignID, clientID string) (bool, error) {
	var clientIsRegistered bool
	res := r.db.WithContext(ctx).Raw("SELECT EXISTS (SELECT 1 FROM clients WHERE client_id = ?)", clientID).Scan(&clientIsRegistered)
	if res.Error != nil {
		return false, res.Error
	}
	if !clientIsRegistered {
		return false, ErrClientNotRegistered
	}
	var result bool
	res = r.db.WithContext(ctx).Raw("SELECT EXISTS (SELECT 1 FROM impressions WHERE client_id = ? AND campaign_id = ?)", clientID, campaignID).Scan(&result)
	if res.Error != nil {
		return false, res.Error
	}
	return result, nil
}
