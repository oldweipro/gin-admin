package request

type EventStreamResult struct {
	Id      string    `json:"id"`
	Object  string    `json:"object"`
	Created uint      `json:"created"`
	Model   string    `json:"model"`
	Choices []Choices `json:"choices"`
}

type Choices struct {
	Delta        Delta  `json:"delta"`
	Index        int    `json:"index"`
	FinishReason string `json:"finish_reason"`
}

type Delta struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}
