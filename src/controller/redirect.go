package controller

import (
	"github.com/dorianneto/url-shortener/src/model"
	"github.com/gin-gonic/gin"
)

type RedirectController struct{}

func (redirect *RedirectController) Index(c *gin.Context) {
	var redirectModel model.Redirect

	if err := c.ShouldBindUri(&redirectModel); err != nil {
		c.JSON(400, gin.H{"message": err})
		return
	}

	c.JSON(200, gin.H{"code": redirectModel.Code})
}
