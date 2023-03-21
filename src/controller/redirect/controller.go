package redirect

import (
	"net/http"

	"github.com/dorianneto/url-shortener/src/controller/redirect/input"
	job "github.com/dorianneto/url-shortener/src/job/redirect"
	"github.com/dorianneto/url-shortener/src/model"
	"github.com/dorianneto/url-shortener/src/queue"
	repository "github.com/dorianneto/url-shortener/src/repository/redirect"
	"github.com/gin-gonic/gin"
)

type RedirectController struct {
	QueueClient queue.QueueClientInterface
	Repository  repository.RedirectRepositoryInterface
}

func (r *RedirectController) Redirect(c *gin.Context) {
	var query input.FindRedirect

	if err := c.ShouldBindUri(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := r.Repository.Find(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.Redirect(http.StatusMovedPermanently, result.Url)
	c.Abort()
}

func (r *RedirectController) Store(c *gin.Context) {
	var payload input.CreateRedirect

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	data, err := model.NewRedirect(payload.Url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	r.QueueClient.Dispatch(&job.CreateRedirectJob{Payload: data})

	c.JSON(http.StatusCreated, gin.H{"data": data})
}
