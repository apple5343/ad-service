package gpthttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"server/pkg/client/yandex-cloud/gpt"
)

func (c *client) GeneratePrompt(prompt gpt.Prompt) (string, error) {
	messages := make([]GptMessage, len(prompt.Messages))
	for i, message := range prompt.Messages {
		messages[i] = GptMessage{
			Role: message.Role,
			Text: message.Content,
		}
	}
	gptRequest := GptRequest{
		ModelUri: c.cfg.ModelUri(),
		Options: GptOptions{
			Temperature: prompt.Temperature,
			MaxTokens:   c.cfg.MaxTokens(),
		},
		Messages: messages,
	}
	body, err := json.Marshal(gptRequest)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", c.cfg.Endpoint(), bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Api-Key "+c.cfg.APIKey())
	req.Header.Set("x-folder-id", c.cfg.Folder())
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result GptResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	fmt.Println(c.cfg.Endpoint())
	fmt.Println(string(body))
	return result.Result.Alternatives[0].Message.Text, nil
}
