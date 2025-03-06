package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type YandexGptConfig interface {
	APIKey() string
	Endpoint() string
	MaxTokens() int
	Folder() string
	ModelUri() string
}

type yandexGptConfig struct {
	KeyField       string `env:"YANDEX_API_KEY"`
	MaxTokensField int    `env:"YANDEX_MAX_TOKENS" env-default:"2000"`
	FolderField    string `env:"YANDEX_FOLDER"`
	ModelField     string `env:"YANDEX_MODEL" env-default:"yandexgpt-lite"`
	EndpointField  string `env:"YANDEX_ENDPOINT" env-default:"https://llm.api.cloud.yandex.net/foundationModels/v1/completion"`
}

func NewYandexGptConfig() (YandexGptConfig, error) {
	cfg := yandexGptConfig{}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (c *yandexGptConfig) APIKey() string {
	return c.KeyField
}

func (c *yandexGptConfig) ModelUri() string {
	return fmt.Sprintf("gpt://%s/%s", c.FolderField, c.ModelField)
}

func (c *yandexGptConfig) MaxTokens() int {
	return c.MaxTokensField
}

func (c *yandexGptConfig) Folder() string {
	return c.FolderField
}

func (c *yandexGptConfig) Endpoint() string {
	return c.EndpointField
}
