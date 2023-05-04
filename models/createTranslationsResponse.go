package models

type CreateTranslationsResponse struct {
	Text string `json:"text,omitempty"`

	Error Error `json:"error,omitempty"`
}
