package controller

import (
	"net/http"

	"github.com/dorianneto/url-shortener/src/input"
	"github.com/dorianneto/url-shortener/src/model"
	"github.com/gin-gonic/gin"
)

type RedirectController struct{}

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

	url := model.NewRedirect(payload.Url, "fmk782ssd")

	c.JSON(http.StatusOK, gin.H{"data": url})
}
