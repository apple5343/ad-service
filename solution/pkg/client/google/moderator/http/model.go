package moderatorhttp

type ModeratorRequest struct {
	Comment    Comment             `json:"comment"`
	Attributes map[string]struct{} `json:"requestedAttributes"`
}

type Comment struct {
	Text string `json:"text"`
}

type ModeratorResponse struct {
	AttributeScores map[string]Attribute `json:"attributeScores"`
}

type Attribute struct {
	SummaryScore SummaryScore `json:"summaryScore"`
}

type SummaryScore struct {
	Score float64 `json:"value"`
	Type  string  `json:"type"`
}
