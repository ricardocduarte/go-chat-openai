package models

type CreateTranslationsRequest struct {
	File           string  `json:"file,omitempty"`
	Model          string  `json:"model,omitempty"`
	Prompt         string  `json:"prompt,omitempty"`
	ResponseFormat string  `json:"response_format,omitempty"`
	Temperature    float64 `json:"temperature,omitempty"`
}
