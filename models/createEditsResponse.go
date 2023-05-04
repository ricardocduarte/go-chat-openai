package models

type CreateEditsResponse struct {
	Object  string `json:"object,omitempty"`
	Created int    `json:"created,omitempty"`
	Choices []struct {
		Text  string `json:"text,omitempty"`
		Index int    `json:"index,omitempty"`
	} `json:"choices,omitempty"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens,omitempty"`
		CompletionTokens int `json:"completion_tokens,omitempty"`
		TotalTokens      int `json:"total_tokens,omitempty"`
	} `json:"usage,omitempty"`

	Error Error `json:"error,omitempty"`
}
