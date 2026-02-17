package services

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kerokume-go/repos"
	"kerokume-go/schemas/contracts"
	"kerokume-go/utils"
)

func LoginService(ctx *gin.Context) {
	var dto contracts.LoginRequest

	if err := ctx.BindJSON(&dto); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	restaurant, err := repos.FindUniqueByEmail(dto, ctx)

	if err != nil{
		utils.SendError(ctx, http.StatusInternalServerError, "error while trying to find a restaurant")
		return 
	}
	
	token, err := JwtServices().GenerateToken(restaurant.ID, ctx)

	response := contracts.LoginResponse{
		Message: "Login successfully",
		Token: token,
	}

	utils.SendSuccessSimple(ctx, "login", response)
}
