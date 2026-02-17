package router

import (
	"github.com/gin-gonic/gin"

	"kerokume-go/middlewares"
	"kerokume-go/services"
)

func restaurantRoutes(group *gin.RouterGroup) {
	group.GET("/restaurants",middlewares.Auth(),services.GetAllRestaurantService)
	group.GET("/restaurants/:id",middlewares.Auth(),services.GetRestaurantServiceGetByID)
	group.POST("/restaurant",services.RestaurantServiceCreate)
	group.POST("/login", services.LoginService)
	group.PUT("/restaurant/:id",middlewares.Auth(),services.UpdateRestaurantService)
	group.DELETE("/restaurant/:id",middlewares.Auth(),services.DeleteRestaurantService)
}
