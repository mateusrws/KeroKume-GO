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

// TODO Configurar o CRUD do Food para o swagger 

// @Summary Create Food
// @Description Create a new food
// @Tags Food
// @Accept json
// @Produce json
// @Param request body contracts.FoodRequest true "Request body"
// @Success 200 {object} contracts.FoodResponse
// @Failure 400 {object} contracts.ErrorResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /foods [post]
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
	utils.SendSuccess(ctx, "create-food")
}


// @Summary Get All Foods
// @Description Get all foods
// @Tags Food
// @Produce json
// @Success 200 {array} contracts.FoodResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /foods [get]
func FoodServiceGetAll(ctx *gin.Context) {
	foods, err := repos.FindAllFood(ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error finding all foods")
		return
	}
	data := make([]interface{}, len(foods))
	for i := range foods {
		data[i] = contracts.FoodResponse{
			Id: foods[i].ID,
			Name: foods[i].Name,
			Description: foods[i].Description,
			Price: foods[i].Price,
			PathImg: foods[i].PathImg,
			FoodCategory: foods[i].FoodCategory,
			IsAvailable: foods[i].IsAvailable,
		}
	}
	utils.SendSuccessArray(ctx, "find-all-foods", data)
}


// @Summary Get Foods by Menu
// @Description Get all foods by menu id
// @Tags Food
// @Produce json
// @Param id path string true "Menu ID"
// @Success 200 {array} contracts.FoodResponse
// @Failure 400 {object} contracts.ErrorResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /menus/{id}/foods [get]
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
		data[i] = contracts.FoodResponse{
			Id: foods[i].ID,
			Name: foods[i].Name,
			Description: foods[i].Description,
			Price: foods[i].Price,
			PathImg: foods[i].PathImg,
			FoodCategory: foods[i].FoodCategory,
			IsAvailable: foods[i].IsAvailable,
		}
	}
	utils.SendSuccessArray(ctx, "find-all-foods-by-menu-id", data)
}


// @Summary Update Food
// @Description Update food by id
// @Tags Food
// @Accept json
// @Produce json
// @Param id path string true "Food ID"
// @Param request body contracts.FoodRequest true "Request body"
// @Success 200 {object} contracts.FoodResponse
// @Failure 400 {object} contracts.ErrorResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /foods/{id} [put]
func FoodServiceUpdate(ctx *gin.Context) {
	logger := config.NewLogger("FOOD UPDATE HANDLER")
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
	
	restaurantId := utils.GetIdFromJwt(ctx)

	err = repos.UpdateFood(restaurantId, foodId, &food, ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error updating food")
		return
	}
	utils.SendSuccess(ctx, "update-food")
}


// @Summary Delete Food
// @Description Delete food by id
// @Tags Food
// @Param id path string true "Food ID"
// @Success 200 "Food deleted successfully"
// @Failure 400 {object} contracts.ErrorResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /foods/{id} [delete]
func FoodServiceDelete(ctx *gin.Context) {
	foodIdStr := ctx.Param("id")
	foodId, err := uuid.Parse(foodIdStr)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "invalid food id")
		return
	}
	err = repos.DeleteFood(foodId, ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error deleting food")
		return
	}
	utils.SendSuccess(ctx, "delete-food")
}