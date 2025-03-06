package moderatorhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"server/pkg/client/google/moderator"
)

var (
	attributes = []string{
		"SEVERE_TOXICITY", "IDENTITY_ATTACK", "INSULT", "PROFANITY", "THREAT",
	}
)

func (c *client) Moderate(text string) (map[string]moderator.Score, error) {
	attributesMap := make(map[string]struct{}, len(attributes))
	for _, attribute := range attributes {
		attributesMap[attribute] = struct{}{}
	}
	moderatorRequest := ModeratorRequest{
		Comment: Comment{
			Text: text,
		},
		Attributes: attributesMap,
	}

	body, err := json.Marshal(moderatorRequest)
	if err != nil {
		return nil, err
	}
	fmt.Println(c.cfg.URL())
	fmt.Println(string(body))

	req, err := http.NewRequest("POST", c.cfg.URL(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var moderatorResponse ModeratorResponse
	err = json.NewDecoder(resp.Body).Decode(&moderatorResponse)
	if err != nil {
		return nil, err
	}
	result := make(map[string]moderator.Score)
	for attribute, score := range moderatorResponse.AttributeScores {
		result[attribute] = moderator.Score{
			Value: score.SummaryScore.Score,
			Type:  score.SummaryScore.Type,
		}
	}
	return result, nil
}
