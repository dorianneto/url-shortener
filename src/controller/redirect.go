package controller

import (
	"net/http"

	"github.com/dorianneto/url-shortener/src/input"
	"github.com/dorianneto/url-shortener/src/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RedirectController struct {
	Validate *validator.Validate
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

	// TODO: move to an use case/service
	url, err := model.NewRedirect(redirect.Validate, payload.Url, "fmk782ssd")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": url})
}
