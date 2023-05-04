package application

import (
	"chat-openai/models"
	"context"
	"encoding/json"
)

const MODERATIONS_URL = "https://api.openai.com/v1/moderations"

func (c *Client) CreateModerationsRaw(ctx context.Context, r models.CreateModerationsRequest) ([]byte, error) {
	return c.Post(ctx, MODERATIONS_URL, r)
}

func (c *Client) CreateModerations(ctx context.Context, r models.CreateModerationsRequest) (response models.CreateModerationsResponse, err error) {
	raw, err := c.CreateModerationsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
