package main

import (
	"log"

	"github.com/dorianneto/url-shortener/src/controller"
	"github.com/dorianneto/url-shortener/src/job"
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
	queueClient := queue.AsynqClientAdapter{}
	queueServer := queue.AsynqServerdapter{}

	queueServer.RegisterWorker(&job.CreateRedirectJob{})

	go queueServer.RunWorkers()

	redirectController := controller.RedirectController{QueueClient: &queueClient}

	router.GET("/:code", redirectController.Index)
	router.POST("/", redirectController.Store)

	router.Run(":8080")
}
