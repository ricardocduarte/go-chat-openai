package application

import (
	"chat-openai/models"
	"context"
	"encoding/json"
)

const EMBEDDINGS_URL = "https://api.openai.com/v1/embeddings"

func (c *Client) CreateEmbeddingsRaw(ctx context.Context, r models.CreateEmbeddingsRequest) ([]byte, error) {
	return c.Post(ctx, EMBEDDINGS_URL, r)
}

func (c *Client) CreateEmbeddings(ctx context.Context, r models.CreateEmbeddingsRequest) (response models.CreateEmbeddingsResponse, err error) {
	raw, err := c.CreateEmbeddingsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
