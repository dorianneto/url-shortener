package main

import (
	"github.com/dorianneto/url-shortener/src/controller"
	asynqclient "github.com/dorianneto/url-shortener/src/queue/asynq"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router := gin.Default()
	validate := *validator.New() // TODO: decouple it
	queue := asynqclient.AsynqClient{}

	redirectController := controller.RedirectController{Validate: &validate, QueueClient: &queue}

	router.GET("/:code", redirectController.Index)
	router.POST("/", redirectController.Store)

	router.Run(":8080")
}
