package models

type RetrieveModelResponse struct {
	ID          string   `json:"id,omitempty"`
	Object      string   `json:"object,omitempty"`
	OwnedBy     string   `json:"owned_by,omitempty"`
	Permissions []string `json:"permissions,omitempty"`

	Error Error `json:"error,omitempty"`
}
