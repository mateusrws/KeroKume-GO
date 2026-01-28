package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func menutRoutes(group *gin.RouterGroup){
	group.GET("/menu", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message":"Teste Menu",
		})
	})
}