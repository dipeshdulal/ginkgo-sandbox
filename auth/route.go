package auth

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	return gin.Default()
}

func RegisterRoutes(router *gin.Engine, controller *Controller) {
	router.GET("/items", controller.GetItems)
}
