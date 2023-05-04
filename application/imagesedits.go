package application

import (
	"chat-openai/models"
	"context"
	"encoding/json"
)

const IMAGES_EDITS_URL = "https://api.openai.com/v1/images/edits"

func (c *Client) CreateImagesEditsRaw(ctx context.Context, r models.CreateImagesEditsRequest) ([]byte, error) {
	return c.Post(ctx, IMAGES_EDITS_URL, r)
}

func (c *Client) CreateImagesEdits(ctx context.Context, r models.CreateImagesEditsRequest) (response models.CreateImagesEditsResponse, err error) {
	raw, err := c.CreateImagesEditsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
