package application

import (
	"chat-openai/models"
	"context"
	"encoding/json"
)

const MODEL_URL = "https://api.openai.com/v1/models/"

func (c *Client) RetrieveModelRaw(ctx context.Context, id string) ([]byte, error) {
	return c.Get(ctx, MODEL_URL+id, nil)
}

func (c *Client) RetrieveModel(ctx context.Context, id string) (response models.RetrieveModelResponse, err error) {
	raw, err := c.RetrieveModelRaw(ctx, id)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
