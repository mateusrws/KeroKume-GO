package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"kerokume-go/repos"
	"kerokume-go/utils"
)

func IfExistInDb(s string, ctx *gin.Context) bool {
	id, err := uuid.Parse(s)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "Error to convert string to uudi")
	}
	_, err = repos.FindUniqueRestaurant(id, ctx)
	if err != nil {
		return false
	}

	return true
}
