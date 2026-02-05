package router

import (
	"github.com/gin-gonic/gin"

	"kerokume-go/services"
)

func menutRoutes(group *gin.RouterGroup) {
	group.GET("/menus", services.GetAllMenuService)
	group.GET("/menus/:id", services.GetMenuServiceGetByRestaurantID)
	group.POST("/menu", services.CreateMenuService)
	group.PUT("/menu/:id", services.UpdateMenuService)
}
