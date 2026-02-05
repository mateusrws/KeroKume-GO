package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"kerokume-go/schemas"
	"kerokume-go/utils"
)

// Create Restaurant

func SaveRestaurant(restaurant schemas.Restaurant, ctx *gin.Context) {
	if err := db.Create(&restaurant).Error; err != nil {
		logger.Errf("error creating restaurant: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error creating restaurant on database")
		return
	}
	utils.SendSuccess(ctx, "create-restaurant", []interface{}{})
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

func UpdateRestaurant(restaurantId uuid.UUID, restaurant *schemas.Restaurant, ctx *gin.Context) error {
	if err := db.Model(&schemas.Restaurant{}).Where("id = ?", restaurantId).Updates(restaurant).Error; err != nil {
		logger.Errf("error updating restaurant: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error updating restaurant")
		return err
	}
	return nil
}