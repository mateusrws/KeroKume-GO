package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "kerokume-go/docs"
	"kerokume-go/repos"

	_ "kerokume-go/docs"
)

func initializeRoutes(router *gin.Engine) {
	// Inittialize Handler
	repos.Init()
	BasePath := "api/v1"
	docs.SwaggerInfo.BasePath = BasePath
	v1 := router.Group(BasePath)
	{
		restaurantRoutes(v1)
		menutRoutes(v1)
		foodRoutes(v1)
	}
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
