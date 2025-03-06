package ai

import "context"

const maxScore = 0.6

func (s *aiService) ModerateText(ctx context.Context, text string) (bool, []string, error) {
	res, err := s.moderatorClient.Moderate(text)
	if err != nil {
		return false, nil, err
	}
	attributes := []string{}
	for atattribute, score := range res {
		if score.Value > maxScore {
			attributes = append(attributes, atattribute)
		}
	}
	if len(attributes) > 0 {
		return false, attributes, nil
	}
	return true, nil, nil
}
