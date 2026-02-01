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
	utils.SendSuccess(ctx, "create-restaurant")
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