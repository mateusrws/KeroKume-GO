package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"kerokume-go/repos"
	"kerokume-go/schemas"
	"kerokume-go/schemas/contracts"
	"kerokume-go/schemas/validates"
	"kerokume-go/utils"
)

// TODO Configurar o CRUD do Menu para o swagger 

// @Summary Create Menu
// @Description Create a new menu for a restaurant
// @Tags Menu
// @Accept json
// @Produce json
// @Param request body contracts.MenuRequest true "Request body"
// @Success 200 {object} contracts.MenuResponse
// @Failure 400 {object} contracts.ErrorResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /menus [post]
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

	utils.SendSuccessSimple(ctx, "create-menu", menu)
}


// @Summary Get All Menus
// @Description Get all menus
// @Tags Menu
// @Produce json
// @Success 200 {array} contracts.MenuResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /menus [get]
func GetAllMenuService(ctx *gin.Context) {
	menus, err := repos.FindAllMenu(ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	data := make([]interface{}, len(menus))
	for i := range menus {
		data[i] = contracts.MenuResponse{
			Id: menus[i].ID,
			Name: menus[i].Name,
		}
	}
	utils.SendSuccessArray(ctx, "find-all-menus", data)
}


// @Summary Get Menus by Restaurant
// @Description Get all menus by restaurant id
// @Tags Menu
// @Produce json
// @Param id path string true "Restaurant ID"
// @Success 200 {array} contracts.MenuResponse
// @Failure 400 {object} contracts.ErrorResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /restaurants/{id}/menus [get]
func GetMenuServiceGetByRestaurantID(ctx *gin.Context) {
	restaurantIdStr := ctx.Param("id")
	restaurantId, err := uuid.Parse(restaurantIdStr)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "invalid restaurant id")
		return
	}
	menus, err := repos.FindAllByRestaurantId(restaurantId, ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error finding all menus by restaurant id")
		return
	}
	data := make([]interface{}, len(menus))
	for i := range menus {
		data[i] = contracts.MenuResponse{
			Id: menus[i].ID,
			Name: menus[i].Name,
		}
	}
	utils.SendSuccessArray(ctx, "find-all-menus-by-restaurant-id", data)
}


// @Summary Update Menu
// @Description Update menu by id
// @Tags Menu
// @Accept json
// @Produce json
// @Param id path string true "Menu ID"
// @Param request body contracts.MenuRequest true "Request body"
// @Success 200 {object} contracts.MenuResponse
// @Failure 400 {object} contracts.ErrorResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /menus/{id} [put]
func UpdateMenuService(ctx *gin.Context) {
	var dto contracts.MenuRequest
	if err := ctx.BindJSON(&dto); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := validates.ValidateCreateMenu(&dto); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	menuIdStr := ctx.Param("id")
	menuId, err := uuid.Parse(menuIdStr)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "invalid menu id")
		return
	}
	restaurantId := utils.GetIdFromJwt(ctx)
	
	menu := schemas.Menu{
		Name:         dto.Name,
		RestaurantId: dto.RestaurantId,
	}
	
	if err := repos.UpdateMenu(restaurantId, menuId, &menu, ctx); err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendSuccess(ctx, "update-menu")
}

// @Summary Delete Menu
// @Description Delete menu by id
// @Tags Menu
// @Param id path string true "Menu ID"
// @Success 200 "Menu deleted successfully"
// @Failure 400 {object} contracts.ErrorResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /menus/{id} [delete]
func DeleteMenuService(ctx *gin.Context) {
	menuIdStr := ctx.Param("id")
	menuId, err := uuid.Parse(menuIdStr)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "invalid menu id")
		return
	}
	err = repos.DeleteMenu(menuId, ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error deleting menu")
		return
	}
	utils.SendSuccess(ctx, "delete-menu")
}