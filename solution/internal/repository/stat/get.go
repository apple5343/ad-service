package stat

import (
	"context"
	"fmt"
	"server/internal/model"
	"server/internal/repository/campaign"
	"server/internal/repository/stat/converter"
	repo "server/internal/repository/stat/model"
	"server/pkg/time"
)

func (r *statRepository) GetByCampaign(ctx context.Context, campaignID string) (*model.Stat, error) {
	stat := repo.Stat{}
	query := `SELECT campaigns.impression_count, campaigns.click_count,
    COALESCE(click_stats.total_click_cost, 0) AS spent_clicks,
    COALESCE(impression_stats.total_impression_cost, 0) AS spent_impressions,
    COALESCE(click_stats.total_click_cost, 0) + COALESCE(impression_stats.total_impression_cost, 0) AS spent_total,
    CASE 
        WHEN campaigns.impression_count = 0 THEN 0 
        ELSE (campaigns.click_count::FLOAT / campaigns.impression_count) * 100 
    END AS conversion
	FROM campaigns
	LEFT JOIN (SELECT campaign_id, SUM(cost) AS total_click_cost FROM clicks GROUP BY campaign_id) 
	click_stats ON click_stats.campaign_id = campaigns.campaign_id
	LEFT JOIN (SELECT campaign_id, SUM(cost) AS total_impression_cost FROM impressions GROUP BY campaign_id) 
	impression_stats ON impression_stats.campaign_id = campaigns.campaign_id WHERE campaigns.campaign_id = ?`
	res := r.db.WithContext(ctx).Raw(query, campaignID).Scan(&stat)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, campaign.ErrCampaignNotFound
	}
	fmt.Println(stat)
	return converter.FromRepoToStat(&stat), nil
}

func (r *statRepository) GetByCampaignDaily(ctx context.Context, campaignID string) (*model.Stat, error) {
	stat := repo.Stat{}
	query := `SELECT impression_stats.impression_count AS impression_count, click_stats.click_count AS click_count,
    COALESCE(click_stats.total_click_cost, 0) AS spent_clicks,
    COALESCE(impression_stats.total_impression_cost, 0) AS spent_impressions,
    COALESCE(click_stats.total_click_cost, 0) + COALESCE(impression_stats.total_impression_cost, 0) AS spent_total,
    CASE 
        WHEN campaigns.impression_count = 0 THEN 0 
        ELSE (click_stats.click_count::FLOAT / impression_stats.impression_count) * 100 
    END AS conversion
	FROM campaigns
	LEFT JOIN (SELECT campaign_id, SUM(cost) AS total_click_cost, COUNT(*) AS click_count FROM clicks WHERE date = ? GROUP BY campaign_id) 
	click_stats ON click_stats.campaign_id = campaigns.campaign_id
	LEFT JOIN (SELECT campaign_id, SUM(cost) AS total_impression_cost, COUNT(*) AS impression_count FROM impressions WHERE date = ? GROUP BY campaign_id) 
	impression_stats ON impression_stats.campaign_id = campaigns.campaign_id WHERE campaigns.campaign_id = ?`
	day := time.Day()
	println(day)
	res := r.db.WithContext(ctx).Raw(query, day, day, campaignID).Scan(&stat)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, campaign.ErrCampaignNotFound
	}

	return converter.FromRepoToStat(&stat), nil
}

func (r *statRepository) GetByAdvertiser(ctx context.Context, advertiserID string) (*model.Stat, error) {
	stat := repo.Stat{}
	query := `WITH st AS (
				SELECT 
					campaigns.campaign_id,
					campaigns.advertiser_id,
					campaigns.impression_count,
					campaigns.click_count,
					COALESCE(click_stats.total_click_cost, 0) AS spent_clicks,
					COALESCE(impression_stats.total_impression_cost, 0) AS spent_impressions,
					COALESCE(click_stats.total_click_cost, 0) + COALESCE(impression_stats.total_impression_cost, 0) AS spent_total
				FROM campaigns
				LEFT JOIN (SELECT campaign_id, SUM(cost) AS total_click_cost FROM clicks GROUP BY campaign_id
				) click_stats ON click_stats.campaign_id = campaigns.campaign_id
				LEFT JOIN (SELECT campaign_id, SUM(cost) AS total_impression_cost FROM impressions GROUP BY campaign_id
				) impression_stats ON impression_stats.campaign_id = campaigns.campaign_id
			)
			SELECT 
				advertisers.advertiser_id, SUM(st.spent_total) as spent_total, SUM(st.spent_clicks) AS spent_clicks, SUM(st.spent_impressions) AS spent_impressions, SUM(st.impression_count) AS impression_count, SUM(st.click_count) AS click_count, CASE WHEN SUM(st.impression_count) = 0 THEN 0 ELSE (SUM(st.click_count)::FLOAT / SUM(st.impression_count)) * 100 END AS conversion
			FROM advertisers
			LEFT JOIN st ON st.advertiser_id = advertisers.advertiser_id
			WHERE advertisers.advertiser_id = ? GROUP BY advertisers.advertiser_id;`

	res := r.db.WithContext(ctx).Raw(query, advertiserID).Scan(&stat)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, campaign.ErrCampaignNotFound
	}

	return converter.FromRepoToStat(&stat), nil
}

func (r *statRepository) GetByAdvertiserDaily(ctx context.Context, advertiserID string) (*model.Stat, error) {
	stat := repo.Stat{}
	query := `WITH st AS (
				SELECT 
					campaigns.campaign_id,
					campaigns.advertiser_id,
					impression_stats.impression_count AS impression_count,
					click_stats.click_count AS click_count,
					COALESCE(click_stats.total_click_cost, 0) AS spent_clicks,
					COALESCE(impression_stats.total_impression_cost, 0) AS spent_impressions,
					COALESCE(click_stats.total_click_cost, 0) + COALESCE(impression_stats.total_impression_cost, 0) AS spent_total
				FROM campaigns
				LEFT JOIN (SELECT campaign_id, COUNT(*) AS click_count, SUM(cost) AS total_click_cost FROM clicks WHERE date = ? GROUP BY campaign_id
				) click_stats ON click_stats.campaign_id = campaigns.campaign_id
				LEFT JOIN (SELECT campaign_id, COUNT(*) AS impression_count, SUM(cost) AS total_impression_cost FROM impressions WHERE date = ? GROUP BY campaign_id
				) impression_stats ON impression_stats.campaign_id = campaigns.campaign_id
			)
			SELECT 
				advertisers.advertiser_id, SUM(st.spent_total) as spent_total, SUM(st.spent_clicks) AS spent_clicks, SUM(st.spent_impressions) AS spent_impressions, SUM(st.impression_count) AS impression_count, SUM(st.click_count) AS click_count, CASE WHEN SUM(st.impression_count) = 0 THEN 0 ELSE (SUM(st.click_count)::FLOAT / SUM(st.impression_count)) * 100 END AS conversion
			FROM advertisers
			LEFT JOIN st ON st.advertiser_id = advertisers.advertiser_id
			WHERE advertisers.advertiser_id = ? GROUP BY advertisers.advertiser_id;`
	day := time.Day()
	res := r.db.WithContext(ctx).Raw(query, day, day, advertiserID).Scan(&stat)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, campaign.ErrCampaignNotFound
	}

	return converter.FromRepoToStat(&stat), nil
}
