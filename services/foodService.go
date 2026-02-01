package services

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kerokume-go/config"
	"kerokume-go/repos"
	"kerokume-go/schemas"
	"kerokume-go/schemas/contracts"
	"kerokume-go/schemas/validates"
	"kerokume-go/utils"
)

func FoodServiceCreate(ctx *gin.Context) {
	logger := config.NewLogger("FOOD CREATE HANDLER")
	var dto contracts.FoodRequest
	if err := ctx.BindJSON(&dto); err != nil {
		utils.SendError(ctx, 400, err.Error())
		return
	}

	if err := validates.ValidateFoodRequest(&dto); err != nil {
		logger.Errf("validation error: %v", err)
		utils.SendError(ctx, 400, err.Error())
		return
	}

	_, err := repos.FindUniqueMenu(dto.MenuId, ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusUnauthorized, "Not exist a Menu with this id")
		return
	}

	if dto.FoodCategory != "COMIDA" && dto.FoodCategory != "BEBIDA" {
		utils.SendError(ctx, http.StatusUnauthorized, "Not is permited this food category")
		return
	}

	food := schemas.Food{
		Name:         dto.Name,
		Description:  dto.Description,
		Price:        dto.Price,
		PathImg:      "",
		FoodCategory: dto.FoodCategory,
		MenuId:       dto.MenuId,
	}

	repos.SaveFood(&food, ctx)
}
