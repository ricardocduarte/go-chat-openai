package application

import (
	"chat-openai/models"
	"context"
	"encoding/json"
)

const EDITS_URL = "https://api.openai.com/v1/edits"

func (c *Client) CreateEditsRaw(ctx context.Context, r models.CreateEditsRequest) ([]byte, error) {
	return c.Post(ctx, EDITS_URL, r)
}

func (c *Client) CreateEdits(ctx context.Context, r models.CreateEditsRequest) (response models.CreateEditsResponse, err error) {
	raw, err := c.CreateEditsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
