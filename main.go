package main

import (
	"github.com/dorianneto/url-shortener/src/controller"
	asynqclient "github.com/dorianneto/url-shortener/src/queue/asynq"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	queue := asynqclient.AsynqClient{}

	redirectController := controller.RedirectController{QueueClient: &queue}

	router.GET("/:code", redirectController.Index)
	router.POST("/", redirectController.Store)

	router.Run(":8080")
}
