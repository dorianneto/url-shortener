package main

import (
	"log"
	"os"

	controller "github.com/dorianneto/url-shortener/src/controller/redirect"
	database "github.com/dorianneto/url-shortener/src/database/firestore"
	job "github.com/dorianneto/url-shortener/src/job/redirect"
	queue "github.com/dorianneto/url-shortener/src/queue/asynq"
	repository "github.com/dorianneto/url-shortener/src/repository/redirect"
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
	redirectRepository := repository.RedirectRepository{
		Database: &database,
	}

	defer database.Close()

	queueServer.RegisterWorker(&job.CreateRedirectJob{
		Repository: &redirectRepository,
	})

	go queueServer.RunWorkers()

	redirectController := controller.RedirectController{
		QueueClient: &queueClient,
		Repository:  &redirectRepository,
	}

	router.GET("/:code", redirectController.Redirect)
	router.POST("/", redirectController.Store)

	router.Run(":8080")
}
