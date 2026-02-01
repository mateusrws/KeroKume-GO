package services

import (
	"github.com/gin-gonic/gin"

	"kerokume-go/repos"
	"kerokume-go/schemas"
	"kerokume-go/schemas/contracts"
	"kerokume-go/schemas/validates"
	"kerokume-go/utils"
)

func RestaurantServiceCreate(ctx *gin.Context) {
	var dto contracts.RestaurantRequest
	if err := ctx.BindJSON(&dto); err != nil {
		utils.SendError(ctx, 400, err.Error())
		return
	}
	if err := validates.ValidateCreateRestaurant(&dto); err != nil {
		logger.Errf("validation error: %v", err)
		utils.SendError(ctx, 400, err.Error())
		return
	}

	passHashed, err := utils.HashPass(dto.Password)
	if err != nil {
		logger.Errf("error hashing password: %v", err)
		utils.SendError(ctx, 500, "Error in Hashing Password")
		return
	}

	restaurant := schemas.Restaurant{
		Name:        dto.Name,
		Description: dto.Description,
		Password:    passHashed,
	}

	repos.SaveRestaurant(restaurant, ctx)
}
