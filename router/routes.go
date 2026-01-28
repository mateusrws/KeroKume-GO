package router

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		restaurantRoutes(v1)
		menutRoutes(v1)
		foodRoutes(v1)
	}
}
