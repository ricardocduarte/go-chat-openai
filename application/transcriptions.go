package application

import (
	"chat-openai/models"
	"context"
	"encoding/json"
)

const TRANSCRIPTIONS_URL = "https://api.openai.com/v1/audio/transcriptions"

func (c *Client) CreateTranscriptionsRaw(ctx context.Context, r models.CreateTranscriptionsRequest) ([]byte, error) {
	return c.Post(ctx, TRANSCRIPTIONS_URL, r)
}

func (c *Client) CreateTranscriptions(ctx context.Context, r models.CreateTranscriptionsRequest) (response models.CreateTranscriptionsResponse, err error) {
	raw, err := c.CreateTranscriptionsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
