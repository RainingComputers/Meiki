package health

import (
	"github.com/gin-gonic/gin"
)

func getHealthStatusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Healthy")
	}
}

func CreateRoutes(router *gin.RouterGroup) {
	router.GET("/", getHealthStatusHandler())
}
