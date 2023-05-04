package models

type CreateImagesEditsRequest struct {
	Image          string `json:"image,omitempty"`
	Mask           string `json:"mask,omitempty"`
	Prompt         string `json:"prompt,omitempty"`
	N              int    `json:"n,omitempty"`
	Size           string `json:"size,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}
