package ai

import (
	"context"
	"server/pkg/client/yandex-cloud/gpt"
)

const prompt = "Напиши описание для рекламы товара исходя из названия рекламодателя и товара. Описание должно быть не слишком большое"
const temperature = 0.5

func (s *aiService) GenerateCampaignDescription(ctx context.Context, adveristerName, campaignName string) (string, error) {
	messages := []gpt.Message{
		{
			Role:    "system",
			Content: prompt,
		},
		{
			Role:    "user",
			Content: adveristerName,
		},
		{
			Role:    "user",
			Content: campaignName,
		},
	}
	description, err := s.gptClient.GeneratePrompt(gpt.Prompt{
		Temperature: temperature,
		Messages:    messages,
	})
	if err != nil {
		return "", err
	}
	return description, nil
}
