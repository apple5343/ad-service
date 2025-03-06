package moderator

type Moderator interface {
	Moderate(text string) (map[string]Score, error)
}

type Score struct {
	Value float64
	Type  string
}
