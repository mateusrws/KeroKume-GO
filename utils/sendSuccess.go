package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccessArray(ctx *gin.Context, op string, dto []interface{}) {
	if len(dto) > 0 {
		ctx.Header("content/type", "application/json")
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("operation from handler: %s successfull", op),
			"data": dto,
		})
		return
	}
	ctx.Header("content/type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
	})
}


func SendSuccessSimple(ctx *gin.Context, op string, dto interface{}) {
	ctx.Header("content/type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data": dto,
	})
}

func SendSuccess(ctx *gin.Context, op string) {
	ctx.Header("content/type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
	})
}

