package gpt

type YandexGptClient interface {
	GeneratePrompt(prompt Prompt) (string, error)
}

type Prompt struct {
	Temperature float64
	Messages    []Message
}

type Message struct {
	Role    string
	Content string
}
