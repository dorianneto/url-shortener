package main

import (
	"log"
	"os"

	controller "github.com/dorianneto/url-shortener/internal/controller/redirect"
	"github.com/dorianneto/url-shortener/internal/database/firestore"
	"github.com/dorianneto/url-shortener/internal/job"
	"github.com/dorianneto/url-shortener/internal/queue/asynq"
	"github.com/dorianneto/url-shortener/internal/repository"
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

	queueClient := asynq.NewAsynqClientAdapter()
	queueServer := asynq.NewAsynqServerdapter()

	database := firestore.NewFirestoreAdapter()
	repository := repository.NewRepository(database)

	defer database.Close()

	job := job.NewCreateRedirectJob(repository)

	queueServer.RegisterWorker(job)

	go queueServer.RunWorkers()

	redirectController := controller.RedirectController{
		QueueClient: queueClient,
		Repository:  repository,
		Job:         job,
	}

	router.GET("/:code", redirectController.Redirect)
	router.POST("/", redirectController.Store)

	router.Run(":8080")
}
