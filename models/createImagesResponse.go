package models

type CreateImagesResponse struct {
	Created int `json:"created,omitempty"`
	Data    []struct {
		URL string `json:"url,omitempty"`
	} `json:"data,omitempty"`

	Error Error `json:"error,omitempty"`
}
