package router

import (
	"github.com/gin-gonic/gin"

	"kerokume-go/services"
)

func restaurantRoutes(group *gin.RouterGroup) {
	group.POST("/restaurant", services.RestaurantServiceCreate)
}
