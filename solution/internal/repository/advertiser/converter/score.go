package converter

import (
	"server/internal/model"
	repo "server/internal/repository/advertiser/model"
)

func ToRepoFromScore(score *model.Score) *repo.Score {
	return &repo.Score{
		AdvertiserID: score.AdvertiserID,
		ClientID:     score.ClientID,
		Score:        score.Score,
	}
}
