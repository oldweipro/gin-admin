package request

// ChatCompletion OpenAI的参数列表
type ChatCompletion struct {
	Model       string        `json:"model"`
	Messages    []ChatMessage `json:"messages"`
	Stream      bool          `json:"stream"`
	Temperature float32       `json:"temperature"`
	TopP        float32       `json:"top_p"`
	//N                int           `json:"n"`
	MaxTokens int `json:"max_tokens"`
	//Stop             string        `json:"stop"`
	PresencePenalty float32 `json:"presence_penalty"`
	//FrequencyPenalty float32       `json:"frequency_penalty"`
	User string `json:"user"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatHistory struct {
	Uuid   uint   `json:"uuid"`
	Title  string `json:"title"`
	IsEdit bool   `json:"isEdit"`
}
