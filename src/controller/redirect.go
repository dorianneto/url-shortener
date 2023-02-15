package controller

import (
	"net/http"

	"github.com/dorianneto/url-shortener/src/input"
	job "github.com/dorianneto/url-shortener/src/job/create_redirect"
	"github.com/dorianneto/url-shortener/src/model"
	"github.com/dorianneto/url-shortener/src/queue"
	"github.com/gin-gonic/gin"
)

type RedirectController struct {
	QueueClient queue.QueueClientInterface
}

func (redirect *RedirectController) Index(c *gin.Context) {
	code := c.Param("code")

	c.JSON(http.StatusOK, gin.H{"code": code})
}

func (redirect *RedirectController) Store(c *gin.Context) {
	var payload input.CreateUrlInput

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	// TODO: move to job
	url, err := model.NewRedirect(payload.Url, "fmk782ssd")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	redirect.QueueClient.Dispatch(&job.CreateRedirectJob{Data: url})

	c.JSON(http.StatusOK, gin.H{"data": url})
}
