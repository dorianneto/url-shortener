package main

import (
	"log"
	"os"

	controller "github.com/dorianneto/url-shortener/src/controller/redirect"
	database "github.com/dorianneto/url-shortener/src/database/firestore"
	job "github.com/dorianneto/url-shortener/src/job/redirect"
	queue "github.com/dorianneto/url-shortener/src/queue/asynq"
	"github.com/dorianneto/url-shortener/src/repository"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "prod" {
		return
	}

	err := godotenv.Load(".env." + appEnv)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := gin.Default()

	queueClient := queue.AsynqClientAdapter{}
	queueServer := queue.AsynqServerdapter{}

	database := database.FirestoreAdapter{}
	repository := repository.NewRepository(&database)

	defer database.Close()

	queueServer.RegisterWorker(&job.CreateRedirectJob{
		Repository: repository,
	})

	go queueServer.RunWorkers()

	redirectController := controller.RedirectController{
		QueueClient: &queueClient,
		Repository:  repository,
	}

	router.GET("/:code", redirectController.Redirect)
	router.POST("/", redirectController.Store)

	router.Run(":8080")
}
