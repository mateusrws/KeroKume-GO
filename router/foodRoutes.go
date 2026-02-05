package router

import (
	"github.com/gin-gonic/gin"

	"kerokume-go/services"
)

func foodRoutes(group *gin.RouterGroup) {
	group.GET("/foods", services.FoodServiceGetAll)
	group.GET("/foods/:id", services.FoodServiceGetByMenuID)
	group.POST("/foods", services.FoodServiceCreate)
	group.PUT("/foods/:id", services.FoodServiceUpdate)
	group.DELETE("/foods/:id", services.FoodServiceDelete)
}
