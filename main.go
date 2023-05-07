package main

import (
	"chat-openai/handlers"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	if os.Getenv("APP_ENV") != "production" {
		log.Println(os.Getenv("APP_ENV"))
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

	}

	router := httprouter.New()
	initRouter(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		ExposedHeaders:   []string{"Accept", "token", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Access-Control-Request-Headers", "Access-Control-Request-Method", "Connection", "Host", "Origin", "User-Agent", "Referer", "Cache-Control", "X-header"},
		AllowedMethods:   []string{"PATCH", "POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8000", handler))
}

func initRouter(router *httprouter.Router) {
	router.POST("/message/send", handlers.CompletionsHandler)
}
