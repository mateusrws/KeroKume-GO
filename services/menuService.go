package services

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"kerokume-go/repos"
	"kerokume-go/schemas"
	"kerokume-go/schemas/contracts"
	"kerokume-go/schemas/validates"
	"kerokume-go/utils"
)

func CreateMenuService(ctx *gin.Context) {
	var dto contracts.MenuRequest

	if err := ctx.BindJSON(&dto); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := validates.ValidateCreateMenu(&dto); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	_, err := repos.FindUniqueRestaurant(dto.RestaurantId, ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusUnauthorized, "Not exist a Restaurant with this id")
		return
	}

	menu := schemas.Menu{
		Name:         dto.Name,
		RestaurantId: dto.RestaurantId,
	}

	if err := repos.SaveMenu(&menu); err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(ctx, "create-menu")
}
