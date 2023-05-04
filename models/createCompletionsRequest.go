package models

type CreateCompletionsRequest struct {
	Model            string            `json:"model,omitempty"`
	Messages         []Message         `json:"messages,omitempty"`
	Prompt           StrArray          `json:"prompt,omitempty"`
	Suffix           string            `json:"suffix,omitempty"`
	MaxTokens        int               `json:"max_tokens,omitempty"`
	Temperature      float64           `json:"temperature,omitempty"`
	TopP             float64           `json:"top_p,omitempty"`
	N                int               `json:"n,omitempty"`
	Stream           bool              `json:"stream,omitempty"`
	LogProbs         int               `json:"logprobs,omitempty"`
	Echo             bool              `json:"echo,omitempty"`
	Stop             StrArray          `json:"stop,omitempty"`
	PresencePenalty  float64           `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64           `json:"frequency_penalty,omitempty"`
	BestOf           int               `json:"best_of,omitempty"`
	LogitBias        map[string]string `json:"logit_bias,omitempty"`
	User             string            `json:"user,omitempty"`
}
