package router

import (
	"github.com/gin-gonic/gin"

	"kerokume-go/repos"
)

func initializeRoutes(router *gin.Engine) {
	// Inittialize Handler
	repos.Init()

	v1 := router.Group("api/v1")
	{
		restaurantRoutes(v1)
		menutRoutes(v1)
		foodRoutes(v1)
	}
}
