package moderatorhttp

import (
	"net/http"
	"server/internal/config"
	"server/pkg/client/google/moderator"
)

type client struct {
	cfg        config.PerspectiveConfig
	httpClient *http.Client
}

func NewClient(cfg config.PerspectiveConfig) moderator.Moderator {
	return &client{cfg: cfg, httpClient: &http.Client{}}
}
