package app

import (
	"server/internal/config"
	adsRepo "server/internal/repository/ads"
	advertiserRepo "server/internal/repository/advertiser"
	campaignRepo "server/internal/repository/campaign"
	campaignObjectMock "server/internal/repository/campaign_object/mock"
	campaignObjectRepo "server/internal/repository/campaign_object/s3"
	clientRepo "server/internal/repository/client"
	statRepo "server/internal/repository/stat"
	timeRepo "server/internal/repository/time"
	adsService "server/internal/service/ads"
	advertiserService "server/internal/service/advertiser"
	"server/internal/service/ai"
	aiService "server/internal/service/ai"
	campaignService "server/internal/service/campaign"
	clientService "server/internal/service/client"
	statService "server/internal/service/stat"
	timeService "server/internal/service/time"
	"server/internal/transport/http"
	moderatorhttp "server/pkg/client/google/moderator/http"
	"server/pkg/client/postgres"
	"server/pkg/client/redis"
	"server/pkg/client/s3"
	gpthttp "server/pkg/client/yandex-cloud/gpt/http"
	"server/pkg/logger"
	prometheus "server/pkg/metrics"

	"go.uber.org/fx"
)

func init() {
	loggerConfig, err := config.NewLoggerConfig()
	if err != nil {
		panic(err)
	}
	logger.Init(loggerConfig)
}

func NewApp() *fx.App {
	return fx.New(
		fx.Options(
			http.NewModule(),
			postgres.NewModule(),
			redis.NewModule(),
			s3.NewModule(),
			prometheus.NewModule(),
			gpthttp.NewModule(),
			moderatorhttp.NewModule(),

			aiService.NewModule(),

			clientRepo.NewModule(),
			clientService.NewModule(),

			advertiserRepo.NewModule(),
			advertiserService.NewModule(),

			campaignObjectRepo.NewModule(),
			campaignRepo.NewModule(),
			campaignService.NewModule(),

			timeRepo.NewModule(),
			timeService.NewModule(),

			adsRepo.NewModule(),
			adsService.NewModule(),

			statRepo.NewModule(),
			statService.NewModule(),
		),
	)
}

func NewTestApp() *fx.App {
	return fx.New(
		fx.Options(
			http.NewModule(),
			postgres.NewModule(),
			redis.NewModule(),
			gpthttp.NewModule(),
			moderatorhttp.NewModule(),

			ai.NewModule(),
			campaignObjectMock.NewModule(),

			clientRepo.NewModule(),
			clientService.NewModule(),

			advertiserRepo.NewModule(),
			advertiserService.NewModule(),

			campaignRepo.NewModule(),
			campaignService.NewModule(),

			timeRepo.NewModule(),
			timeService.NewModule(),

			adsRepo.NewModule(),
			adsService.NewModule(),

			statRepo.NewModule(),
			statService.NewModule(),
		),
	)
}
