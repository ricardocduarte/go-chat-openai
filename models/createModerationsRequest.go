package models

type CreateModerationsRequest struct {
	Input StrArray `json:"input,omitempty"`
	Model string   `json:"model,omitempty"`
}
