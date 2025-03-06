package ai

import (
	"server/internal/service"
	"server/pkg/client/google/moderator"
	"server/pkg/client/yandex-cloud/gpt"
)

type aiService struct {
	gptClient       gpt.YandexGptClient
	moderatorClient moderator.Moderator
}

func NewAiService(gptClient gpt.YandexGptClient, moderatorClient moderator.Moderator) service.AiService {
	return &aiService{gptClient: gptClient, moderatorClient: moderatorClient}
}
