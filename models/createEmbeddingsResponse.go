package models

type CreateEmbeddingsResponse struct {
	Object string `json:"object,omitempty"`
	Data   []struct {
		Object    string    `json:"object,omitempty"`
		Embedding []float64 `json:"embedding,omitempty"`
		Index     int       `json:"index,omitempty"`
	} `json:"data,omitempty"`
	Model string `json:"model,omitempty"`
	Usage struct {
		PromptTokens int `json:"prompt_tokens,omitempty"`
		TotalTokens  int `json:"total_tokens,omitempty"`
	} `json:"usage,omitempty"`

	Error Error `json:"error"`
}
