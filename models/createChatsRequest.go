package models

type CreateChatsRequest struct {
	Model            string            `json:"model,omitempty"`
	Messages         []Message         `json:"messages,omitempty"`
	Temperature      float64           `json:"temperature,omitempty"`
	TopP             float64           `json:"top_p,omitempty"`
	N                int               `json:"n,omitempty"`
	Stream           bool              `json:"stream,omitempty"`
	Stop             StrArray          `json:"stop,omitempty"`
	MaxTokens        int               `json:"max_tokens,omitempty"`
	PresencePenalty  float64           `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64           `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]string `json:"logit_bias,omitempty"`
	User             string            `json:"user,omitempty"`
}
