package gpthttp

import (
	"net/http"
	"server/internal/config"
	"server/pkg/client/yandex-cloud/gpt"
)

type client struct {
	cfg        config.YandexGptConfig
	httpClient *http.Client
}

func NewClient(cfg config.YandexGptConfig) gpt.YandexGptClient {
	return &client{cfg: cfg, httpClient: &http.Client{}}
}
