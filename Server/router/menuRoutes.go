package router

import (
	"github.com/gin-gonic/gin"

	"kerokume-go/middlewares"
	"kerokume-go/services"
)

func menutRoutes(group *gin.RouterGroup) {
	group.GET("/menus", middlewares.Auth(),services.GetAllMenuService)
	group.GET("/menus/:id", middlewares.Auth(),services.GetMenuServiceGetByRestaurantID)
	group.POST("/menu", middlewares.Auth(),services.CreateMenuService)
	group.PUT("/menu/:id", middlewares.Auth(),services.UpdateMenuService)
	group.PATCH("/menu/:menu-id", middlewares.Auth(),services.AlterMenuActive)
	group.DELETE("/menu/:id", middlewares.Auth(),services.DeleteMenuService)
}
