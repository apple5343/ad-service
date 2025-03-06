package converter

import (
	"server/internal/model"
	"server/internal/transport/http/converter"
	req "server/internal/transport/http/handlers/advertiser/model"
)

func ToScoreFromReq(score *req.Score) *model.Score {
	return &model.Score{
		ClientID:     converter.FromStringPtr(score.ClientID),
		AdvertiserID: converter.FromStringPtr(score.AdvertiserID),
		Score:        converter.FromIntPtr(score.Score),
	}
}
