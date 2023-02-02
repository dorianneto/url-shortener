package main

import (
	"github.com/dorianneto/url-shortener/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	redirectController := controller.RedirectController{}

	router.GET("/:code", redirectController.Index)
	router.POST("/", redirectController.Store)

	router.Run(":8080")
}
