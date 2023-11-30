package auth

import (
	"ginkgo-tests/contracts"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service contracts.AuthService
}

func NewController(service contracts.AuthService) *Controller {
	return &Controller{
		service: service,
	}
}

func (c Controller) GetItems(ctx *gin.Context) {
	items, err := c.service.GetItems()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "pong",
		"items":   items,
	})
}
