package http

import (
	"context"
	"server/internal/config"
	"server/internal/service"
	"server/internal/transport/http/handlers/ads"
	"server/internal/transport/http/handlers/advertiser"
	"server/internal/transport/http/handlers/campaign"
	"server/internal/transport/http/handlers/client"
	"server/internal/transport/http/handlers/stat"
	"server/internal/transport/http/handlers/time"
	myMiddleware "server/internal/transport/http/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	e   *echo.Echo
	cfg config.HTTPConfig

	client     *client.Handler
	advertiser *advertiser.Handler
	campaign   *campaign.Handler
	ads        *ads.Handler
	stat       *stat.Handler
	time       *time.Handler
}

func NewServer(cfg config.HTTPConfig, clientService service.ClientService,
	advertiserService service.AdvertiserService, campaignService service.CampaignService,
	timeService service.TimeService, adsService service.AdsService,
	statService service.StatService) *Server {

	client := client.NewHandler(clientService)
	advertiser := advertiser.NewHandler(advertiserService)
	campaign := campaign.NewHandler(campaignService)
	ads := ads.NewHandler(adsService)
	time := time.NewHandler(timeService)
	stat := stat.NewHandler(statService)

	return &Server{
		e:   echo.New(),
		cfg: cfg,

		client:     client,
		advertiser: advertiser,
		campaign:   campaign,
		ads:        ads,
		stat:       stat,
		time:       time,
	}
}

func (s *Server) Run() error {
	s.routes()
	return s.e.Start(s.cfg.Address())
}

func (s *Server) Stop(ctx context.Context) error {
	return s.e.Shutdown(ctx)
}

func (s *Server) routes() {
	s.e.Use(middleware.Recover())
	s.e.Use(myMiddleware.MetricsMiddleware)
	s.e.Use(myMiddleware.ErrorHandlerMiddleware)
	s.e.Validator = &CustomValidator{}

	s.e.POST("/ml-scores", s.advertiser.AddScore())
	s.e.POST("/moderate", s.campaign.Moderate())

	clients := s.e.Group("/clients")
	{
		clients.POST("/bulk", s.client.Save())
		clients.GET("/:clientId", s.client.Get())
	}

	advertisers := s.e.Group("/advertisers")
	{
		advertisers.POST("/bulk", s.advertiser.Save())
		advertisers.GET("/:advertiserId", s.advertiser.Get())

		campaigns := advertisers.Group("/:advertiserId/campaigns")
		{
			campaigns.POST("/generate", s.advertiser.Generate())
			campaigns.POST("", s.campaign.Create())
			campaigns.GET("/:campaignId", s.campaign.Get())
			campaigns.POST("/:campaignId/image", s.campaign.SaveImage())
			campaigns.GET("", s.campaign.List())
			campaigns.DELETE("/:campaignId", s.campaign.Delete())
			campaigns.PUT("/:campaignId", s.campaign.Update())
		}
	}

	ads := s.e.Group("/ads")
	{
		ads.GET("", s.ads.Get())
		ads.POST("/:adId/click", s.ads.Click())
	}

	stats := s.e.Group("/stats")
	{
		stats.GET("/campaigns/:campaignId", s.stat.GetByCampaign())
		stats.GET("/campaigns/:campaignId/daily", s.stat.GetByCampaignDaily())
		stats.GET("/advertisers/:advertiserId/campaigns", s.stat.GetByAdvertiser())
		stats.GET("/advertisers/:advertiserId/campaigns/daily", s.stat.GetByAdvertiserDaily())
	}

	s.e.POST("/time/advance", s.time.Set())
	s.e.GET("/time", s.time.Get())
}
