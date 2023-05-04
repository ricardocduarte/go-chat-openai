package models

type ListModelsResponse struct {
	Data []struct {
		ID          string   `json:"id,omitempty"`
		Object      string   `json:"object,omitempty"`
		OwnedBy     string   `json:"owned_by,omitempty"`
		Permissions []string `json:"permissions,omitempty"`
	} `json:"data,omitempty"`
	Object string `json:"object,omitempty"`

	Error Error `json:"error,omitempty"`
}
