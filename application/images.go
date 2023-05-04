package application

import (
	"chat-openai/models"
	"context"
	"encoding/json"
)

const IMAGES_URL = "https://api.openai.com/v1/images/generations"

func (c *Client) CreateImagesRaw(ctx context.Context, r models.CreateImagesRequest) ([]byte, error) {
	return c.Post(ctx, IMAGES_URL, r)
}

func (c *Client) CreateImages(ctx context.Context, r models.CreateImagesRequest) (response models.CreateImagesResponse, err error) {
	raw, err := c.CreateImagesRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
