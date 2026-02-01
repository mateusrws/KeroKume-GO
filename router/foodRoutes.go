package router

import (
	"github.com/gin-gonic/gin"

	"kerokume-go/services"
)

func foodRoutes(group *gin.RouterGroup) {
	group.POST("/foods", services.FoodServiceCreate)
}
