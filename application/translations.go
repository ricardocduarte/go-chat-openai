package application

import (
	"chat-openai/models"
	"context"
	"encoding/json"
)

const TRANSLATIONS_URL = "https://api.openai.com/v1/audio/translations"

func (c *Client) CreateTranslationsRaw(ctx context.Context, r models.CreateTranslationsRequest) ([]byte, error) {
	return c.Post(ctx, TRANSLATIONS_URL, r)
}

func (c *Client) CreateTranslations(ctx context.Context, r models.CreateTranslationsRequest) (response models.CreateTranslationsResponse, err error) {
	raw, err := c.CreateTranslationsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
