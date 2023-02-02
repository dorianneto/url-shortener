package controller

import (
	"net/http"

	"github.com/dorianneto/url-shortener/src/model"
	"github.com/gin-gonic/gin"
)

type RedirectController struct{}

func (redirect *RedirectController) Index(c *gin.Context) {
	code := c.Param("code")

	c.JSON(200, gin.H{"code": code})
}

func (redirect *RedirectController) Store(c *gin.Context) {
	var redirectModel model.Redirect

	if err := c.ShouldBindJSON(&redirectModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.JSON(200, gin.H{"code": redirectModel.Code})
}
