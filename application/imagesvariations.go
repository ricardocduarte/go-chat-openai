package application

import (
	"chat-openai/models"
	"context"
	"encoding/json"
)

const IMAGES_VARIATIONS_URL = "https://api.openai.com/v1/images/variations"

func (c *Client) CreateImagesVariationsRaw(ctx context.Context, r models.CreateImagesVariationsRequest) ([]byte, error) {
	return c.Post(ctx, IMAGES_VARIATIONS_URL, r)
}

func (c *Client) CreateImagesVariations(ctx context.Context, r models.CreateImagesVariationsRequest) (response models.CreateImagesVariationsResponse, err error) {
	raw, err := c.CreateImagesVariationsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
