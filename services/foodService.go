package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

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

func FoodServiceGetAll(ctx *gin.Context) {
	foods, err := repos.FindAllFood(ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error finding all foods")
		return
	}
	data := make([]interface{}, len(foods))
	for i := range foods {
		data[i] = foods[i]
	}
	utils.SendSuccess(ctx, "find-all-foods", data)
}

func FoodServiceGetByMenuID(ctx *gin.Context) {
	menuIdStr := ctx.Param("id")
	menuId, err := uuid.Parse(menuIdStr)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "invalid menu id")
		return
	}
	foods, err := repos.FindAllFoodByMenuId(menuId, ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error finding all foods by menu id")
		return
	}
	data := make([]interface{}, len(foods))
	for i := range foods {
		data[i] = foods[i]
	}
	utils.SendSuccess(ctx, "find-all-foods-by-menu-id", data)
}

func FoodServiceUpdate(ctx *gin.Context) {
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

	foodIdStr := ctx.Param("id")
	foodId, err := uuid.Parse(foodIdStr)

	food := schemas.Food{
		Name:         dto.Name,
		Description:  dto.Description,
		Price:        dto.Price,
		PathImg:      "",
		FoodCategory: dto.FoodCategory,
		MenuId:       dto.MenuId,
	}

	err = repos.UpdateFood(foodId, &food, ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error updating food")
		return
	}
	utils.SendSuccess(ctx, "update-food", []interface{}{})
}
