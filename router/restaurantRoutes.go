package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func restaurantRoutes(group *gin.RouterGroup){
	group.GET("/restaurant", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":"Teste Restaurante",
		})
	})
}