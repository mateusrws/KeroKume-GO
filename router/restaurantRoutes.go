package router

import (
	"github.com/gin-gonic/gin"

	"kerokume-go/services"
)

func restaurantRoutes(group *gin.RouterGroup) {
	group.GET("/restaurants", services.GetAllRestaurantService)
	group.GET("/restaurants/:id", services.GetRestaurantServiceGetByID)
	group.POST("/restaurant", services.RestaurantServiceCreate)
	group.PUT("/restaurant/:id", services.UpdateRestaurantService)
	group.DELETE("/restaurant/:id", services.DeleteRestaurantService)
}
