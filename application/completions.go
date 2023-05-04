package application

import (
	"chat-openai/models"
	"context"
	"encoding/json"
)

const COMPLETIONS_URL = "https://api.openai.com/v1/chat/completions"

func (c *Client) CreateCompletionsRaw(ctx context.Context, r models.CreateCompletionsRequest) ([]byte, error) {
	return c.Post(ctx, COMPLETIONS_URL, r)
}

func (c *Client) CreateCompletions(ctx context.Context, r models.CreateCompletionsRequest) (response models.CreateCompletionsResponse, err error) {
	raw, err := c.CreateCompletionsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
