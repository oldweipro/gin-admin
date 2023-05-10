package example

type ChatGPT struct {
	Options       Options `json:"options"`
	Prompt        string  `json:"prompt"`
	SystemMessage string  `json:"systemMessage"`
	Temperature   int     `json:"temperature"`
	TopP          int     `json:"topP"`
}

type Options struct {
	ConversationId  string `json:"conversationId"`
	ParentMessageId string `json:"parentMessageId"`
}
