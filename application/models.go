package application

import (
	"chat-openai/models"
	"context"
	"encoding/json"
)

const MODELS_URL = "https://api.openai.com/v1/models"

func (c *Client) ListModelsRaw(ctx context.Context) ([]byte, error) {
	return c.Get(ctx, MODELS_URL, nil)
}

func (c *Client) ListModels(ctx context.Context) (response models.ListModelsResponse, err error) {
	raw, err := c.ListModelsRaw(ctx)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
