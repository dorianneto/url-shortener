package controller

import (
	"net/http"

	"github.com/dorianneto/url-shortener/src/input"
	"github.com/dorianneto/url-shortener/src/job"
	"github.com/dorianneto/url-shortener/src/model"
	"github.com/dorianneto/url-shortener/src/queue"
	"github.com/dorianneto/url-shortener/src/repository"
	"github.com/gin-gonic/gin"
)

type RedirectController struct {
	QueueClient queue.QueueClientInterface
	Repository  repository.RepositoryInterface
}

func (r *RedirectController) Index(c *gin.Context) {
	data, err := r.Repository.Find(c.Param("code"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	redirect := data.(*model.Redirect)

	c.Redirect(http.StatusMovedPermanently, redirect.Url)
}

func (redirect *RedirectController) Store(c *gin.Context) {
	var payload input.CreateUrlInput

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	data, err := model.NewRedirect(payload.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	redirect.QueueClient.Dispatch(&job.CreateRedirectJob{Payload: data})

	c.JSON(http.StatusCreated, gin.H{"data": data})
}
