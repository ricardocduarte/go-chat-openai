package models

type CreateTranscriptionsResponse struct {
	Text string `json:"text,omitempty"`

	Error Error `json:"error,omitempty"`
}
