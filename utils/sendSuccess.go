package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccess(ctx *gin.Context, op string, food []interface{}) {
	if len(food) > 0 {
		ctx.Header("content/type", "application/json")
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("operation from handler: %s successfull", op),
			"data": food,
		})
		return
	}
	ctx.Header("content/type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
	})
}
