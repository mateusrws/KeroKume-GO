package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"kerokume-go/schemas"
	"kerokume-go/utils"
)

// Create Food
func SaveFood(food *schemas.Food, ctx *gin.Context) {
	if err := db.Create(food).Error; err != nil {
		logger.Errf("error creating food: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error creating food on database")
		return
	}

	utils.SendSuccess(ctx, "create-food")
}

// Find Food by ID
func FindUniqueFood(id uuid.UUID, ctx *gin.Context) (*schemas.Food, error) {
	var food schemas.Food

	if err := db.First(&food, "id = ?", id).Error; err != nil {
		logger.Errf("error finding food: %v", err)
		utils.SendError(ctx, http.StatusNotFound, "food not found")
		return nil, err
	}

	return &food, nil
}
