package converter

import (
	"server/internal/model"
	repo "server/internal/repository/stat/model"
)

func FromRepoToStat(stat *repo.Stat) *model.Stat {
	return &model.Stat{
		ImpressionsCount: stat.ImpressionCount,
		ClicksCount:      stat.ClickCount,
		Conversion:       stat.Conversion,
		SpentImpressions: stat.SpentImpressions,
		SpentClicks:      stat.SpentClicks,
		SpentTotal:       stat.SpentTotal,
	}
}
