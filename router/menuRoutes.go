package router

import (
	"github.com/gin-gonic/gin"

	"kerokume-go/services"
)

func menutRoutes(group *gin.RouterGroup) {
	group.POST("/menu", services.CreateMenuService)
}
