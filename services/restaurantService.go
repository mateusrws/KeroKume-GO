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

// @BasePath /api/v1

// @Summary Create Restaurant 
// @Description Create a new restaurant
// @Tags Restaurant 
// @Accept json
// @Produce json
// @Param request body contracts.RestaurantRequest true "Request body"
// @Success 200 {object} contracts.RestaurantResponse
// @Failure 400 {object} contracts.ErrorResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /restaurant [post]
func RestaurantServiceCreate(ctx *gin.Context) {
	logger := config.NewLogger("RESTAURANT CREATE HANDLER")
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

	utils.SendSuccess(ctx, "create-restaurant")
}

// TODO Adicionar a configuração para o swagger para o GetAllRestaurantService

// @Summary Get Restaurants
// @Description Get All Restaurants
// @Tags Restaurant
// Accept json
// @Produce json
// @Success 200 {array} contracts.RestaurantResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /restaurants [get]
func GetAllRestaurantService(ctx *gin.Context) {
	restaurants, err := repos.FindAllRestaurant(ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error finding all restaurants")
		return
	}
	data := make([]interface{}, len(restaurants))
	for i := range restaurants{
		data[i] = contracts.RestaurantResponse{
			Id: restaurants[i].ID,
			Name: restaurants[i].Name,
			Email: restaurants[i].Email,
		}
	}
	
	utils.SendSuccessArray(ctx, "find-all-restaurants", []interface{}{restaurants})
}

// TODO Adicionar a configuração para o swagger para o GerRestaurantServiceGetById
// @Summary Get Restaurants
// @Description Get All Restaurants
// @Tags Restaurant
// Accept json
// @Produce json
// @Success 200 {array} contracts.RestaurantResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /restaurants [get]
func GetRestaurantServiceGetByID(ctx *gin.Context) {
	restaurantId := utils.GetIdFromJwt(ctx)
	restaurants, err := repos.FindUniqueRestaurant(restaurantId, ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error finding restaurant by id")
		return
	}
	utils.SendSuccessArray(ctx, "update-restaurant", []interface{}{restaurants})
}

// TODO Adicionar a configuração para o swagger para o UpdateRestaurantService
// @Summary Update Restaurant 
// @Description Update restaurant
// @Tags Restaurant 
// @Success 200 {object} contracts.RestaurantResponse
// @Failure 400 {object} contracts.ErrorResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /restaurant [put]
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
	restaurantId := utils.GetIdFromJwt(ctx)

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

	utils.SendSuccess(ctx, "update-restaurant")
}


// TODO Adicionar a configuração para o swagger para o DeleteRestaurantService
// @Summary Delete Restaurant 
// @Description Delete restaurant
// @Tags Restaurant 
// @Success 200 "Restaurant deleted successfully"
// @Failure 400 {object} contracts.ErrorResponse
// @Failure 500 {object} contracts.ErrorResponse
// @Router /restaurant [delete]
func DeleteRestaurantService(ctx *gin.Context){
	restaurantId := utils.GetIdFromJwt(ctx)
	err := repos.DeleteRestaurant(restaurantId, ctx)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error deleting restaurant")
		return
	}
	utils.SendSuccess(ctx, "delete-restaurant") 
}