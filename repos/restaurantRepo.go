package repos

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"kerokume-go/schemas"
	"kerokume-go/schemas/contracts"
	"kerokume-go/utils"
)

// Create Restaurant

func SaveRestaurant(restaurant schemas.Restaurant, ctx *gin.Context) {
	if err := db.Create(&restaurant).Error; err != nil {
		logger.Errf("error creating restaurant: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error creating restaurant on database")
		return
	}
	utils.SendSuccessArray(ctx, "create-restaurant", []interface{}{})
}

func FindUniqueRestaurant(id uuid.UUID, ctx *gin.Context) (schemas.Restaurant, error) {
	var restaurant schemas.Restaurant
	if err := db.First(&restaurant, id).Error; err != nil {
		logger.Errf("error finding restaurant: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error finding restaurant on database")
		return schemas.Restaurant{}, err
	}
	return restaurant, nil
}

func FindAllRestaurant(ctx *gin.Context) ([]schemas.Restaurant, error) {
	var restaurants []schemas.Restaurant
	if err := db.Find(&restaurants).Error; err != nil {
		logger.Errf("error finding all restaurants: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error finding all restaurants")
		return nil, err
	}
	return restaurants, nil
}

func FindUniqueByEmail(dto contracts.LoginRequest, ctx *gin.Context) (*schemas.Restaurant, error) {
	var restaurant schemas.Restaurant

	err := db.Where("email = ?", dto.Email).First(&restaurant).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.SendError(ctx, http.StatusUnauthorized, "invalid email or password")
			return nil, err
		}

		utils.SendError(ctx, http.StatusInternalServerError, "invalid email or password")
		return nil, err
	}

	return &restaurant, nil
}



func UpdateRestaurant(restaurantId uuid.UUID, restaurant *schemas.Restaurant, ctx *gin.Context) error {
	if err := db.Model(&schemas.Restaurant{}).Where("id = ?", restaurantId).Updates(restaurant).Error; err != nil {
		logger.Errf("error updating restaurant: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error updating restaurant")
		return err
	}
	return nil
}

func DeleteRestaurant(restaurantId uuid.UUID, ctx *gin.Context) error {
	if err := db.Delete(&schemas.Restaurant{}, restaurantId).Error; err != nil {
		logger.Errf("error deleting restaurant: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error deleting restaurant")
		return err
	}
	return nil
}