package application

import (
	"chat-openai/models"
	"context"
	"encoding/json"
)

const CHATS_URL = "https://api.openai.com/v1/chat/completions"

func (c *Client) CreateChatsRaw(ctx context.Context, r models.CreateChatsRequest) ([]byte, error) {
	return c.Post(ctx, CHATS_URL, r)
}

func (c *Client) CreateChats(ctx context.Context, r models.CreateChatsRequest) (response models.CreateChatsResponse, err error) {
	raw, err := c.CreateChatsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
