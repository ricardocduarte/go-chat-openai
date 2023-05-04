package models

type CreateEmbeddingsRequest struct {
	Model string   `json:"model,omitempty"`
	Input StrArray `json:"input,omitempty"`
	User  string   `json:"user,omitempty"`
}
