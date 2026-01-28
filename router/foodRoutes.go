package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func foodRoutes(group *gin.RouterGroup){
	group.GET("/foods", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":"Teste Food",
		})
	})
}