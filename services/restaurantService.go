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

func GetAllRestaurantService(ctx *gin.Context) {
	restaurants, err := repos.FindAllRestaurant(ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error finding all restaurants")
		return
	}
	utils.SendSuccess(ctx, "find-all-restaurants", []interface{}{restaurants})
}

func GetRestaurantServiceGetByID(ctx *gin.Context) {
	restaurantIdStr := ctx.Param("id")
	restaurantId, err := uuid.Parse(restaurantIdStr)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "invalid restaurant id")
		return
	}
	restaurants, err := repos.FindUniqueRestaurant(restaurantId, ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error finding restaurant by id")
		return
	}
	utils.SendSuccess(ctx, "find-restaurant-by-id", []interface{}{restaurants})
}

func UpdateRestaurantService(ctx *gin.Context) {
	var dto contracts.RestaurantRequest
	if err := ctx.BindJSON(&dto); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := validates.ValidateUpdateRestaurant(&dto); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	restaurantIdStr := ctx.Param("id")
	restaurantId, err := uuid.Parse(restaurantIdStr)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, "invalid restaurant id")
		return
	}

	restaurant, err := repos.FindUniqueRestaurant(restaurantId, ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error finding restaurant by id")
		return
	}

	restaurant = schemas.Restaurant{
		Name:        dto.Name,
		Description: dto.Description,
		Password:    restaurant.Password,
	}

	if err := repos.UpdateRestaurant(restaurantId, &restaurant, ctx); err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendSuccess(ctx, "update-restaurant", []interface{}{})
}


func DeleteRestaurantService(ctx *gin.Context){
	restaurantIdStr := ctx.Param("id")
	restaurantId, err := uuid.Parse(restaurantIdStr)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "invalid menu id")
		return
	}
	err = repos.DeleteRestaurant(restaurantId, ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error deleting restaurant")
		return
	}
	utils.SendSuccess(ctx, "delete-restaurant", []interface{}{})
}