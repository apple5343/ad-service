package gpthttp

type GptRequest struct {
	ModelUri string       `json:"modelUri"`
	Options  GptOptions   `json:"completionOptions"`
	Messages []GptMessage `json:"messages"`
}

type GptOptions struct {
	Stream      bool    `json:"stream"`
	Temperature float64 `json:"temperature"`
	MaxTokens   int     `json:"maxTokens"`
}

type GptMessage struct {
	Role string `json:"role"`
	Text string `json:"text"`
}

type GptResponse struct {
	Result GptResult `json:"result"`
}

type GptResult struct {
	Alternatives []GptAlternative `json:"alternatives"`
	Usage        GptUsage         `json:"usage"`
}

type GptAlternative struct {
	Message GptMessage `json:"message"`
	Status  string     `json:"status"`
}

type GptUsage struct {
	InputTextTokens  string `json:"inputTextTokens"`
	CompletionTokens string `json:"completionTokens"`
	TotalTokens      string `json:"totalTokens"`
}
