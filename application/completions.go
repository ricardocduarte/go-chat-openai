package application

import (
	"chat-openai/models"
	"context"
	"encoding/json"
	"fmt"
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

func GetCompletions(client *Client, content string) (response *models.CompletionsResponse, err error) {
	r := models.CreateCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []models.Message{
			{
				Role:    "user",
				Content: content,
			},
		},
		Temperature: 0.7,
	}
	completions, err := client.CreateCompletionsRaw(context.Background(), r)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(completions, &response)
	return response, err

}

func PrintCompletions(client *Client, content string) {
	r := models.CreateCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []models.Message{
			{
				Role:    "user",
				Content: content,
			},
		},
		Temperature: 0.7,
	}
	completions, err := client.CreateCompletionsRaw(context.Background(), r)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(completions))

}
