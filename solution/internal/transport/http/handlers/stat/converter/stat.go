package converter

import (
	"server/internal/model"
	req "server/internal/transport/http/handlers/stat/model"
)

func FromStatToResp(stat *model.Stat) *req.Stat {
	return &req.Stat{
		ImpressionsCount: stat.ImpressionsCount,
		ClicksCount:      stat.ClicksCount,
		Conversion:       stat.Conversion,
		SpentImpressions: stat.SpentImpressions,
		SpentClicks:      stat.SpentClicks,
		SpentTotal:       stat.SpentTotal,
	}
}
