package main

import (
	"log"

	"github.com/dorianneto/url-shortener/src/controller"
	queue "github.com/dorianneto/url-shortener/src/queue/asynq"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()
	queue := queue.AsynqClientAdapter{}

	redirectController := controller.RedirectController{QueueClient: &queue}

	router.GET("/:code", redirectController.Index)
	router.POST("/", redirectController.Store)

	router.Run(":8080")
}
