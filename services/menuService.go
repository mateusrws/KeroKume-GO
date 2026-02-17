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

	menu := schemas.Menu{
		Name:         dto.Name,
		RestaurantId: dto.RestaurantId,
	}

	if err := repos.UpdateMenu(menuId, &menu, ctx); err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendSuccess(ctx, "update-menu")
}

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