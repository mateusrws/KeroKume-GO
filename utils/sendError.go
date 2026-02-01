package utils

import "github.com/gin-gonic/gin"


func SendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("content/type","application/json")
	ctx.JSON(code, gin.H{
		"message": msg,
	})
}