package main

import (
	"chat-openai/application"
	"chat-openai/models"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	client := getClient()

	writeModels(client)

	printCompletions(client)
}

func getClient() *application.Client {
	apiKey := os.Getenv("API_KEY")
	organization := os.Getenv("API_ORG")
	return application.NewClient(apiKey, organization)
}

func writeModels(client *application.Client) {
	models, err := client.ListModelsRaw(context.Background())
	if err != nil {
		panic(err)
	}
	f, err := os.Create("openAi-models-2.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write models to file
	_, err = f.Write(models)
	if err != nil {
		panic(err)
	}
}

func printCompletions(client *application.Client) {
	r := models.CreateCompletionsRequest{
		Model: "gpt-3.5-turbo",
		Messages: []models.Message{
			{
				Role:    "user",
				Content: "criar uma legenda para melinia software crm",
			},
		},
		Temperature: 0.7,
	}
	//O caminho para uma gestão eficiente e uma experiência do cliente excepcional
	completions, err := client.CreateCompletionsRaw(context.Background(), r)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(completions))
	/*
		{
		  "id": "chatcmpl-xxx",
		  "object": "chat.completion",
		  "created": 1678667132,
		  "model": "gpt-3.5-turbo-0301",
		  "usage": {
		    "prompt_tokens": 13,
		    "completion_tokens": 7,
		    "total_tokens": 20
		  },
		  "choices": [
		    {
		      "message": {
		        "role": "assistant",
		        "content": "\n\nThis is a test!"
		      },
		      "finish_reason": "stop",
		      "index": 0
		    }
		  ]
		}
	*/
}
