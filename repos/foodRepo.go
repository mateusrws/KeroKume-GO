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

	utils.SendSuccess(ctx, "create-food", []interface{}{})
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

func FindAllFood(ctx *gin.Context) ([]schemas.Food, error) {
	var foods []schemas.Food

	if err := db.Find(&foods).Error; err != nil {
		logger.Errf("error finding all foods: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error finding all foods")
		return nil, err
	}

	return foods, nil
}

func FindAllFoodByMenuId(menuId uuid.UUID, ctx *gin.Context) ([]schemas.Food, error) {
	var foods []schemas.Food

	if err := db.Where("menu_id = ?", menuId).Find(&foods).Error; err != nil {
		logger.Errf("error finding all foods by menu id: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error finding all foods by menu id")
		return nil, err
	}

	return foods, nil
}


func UpdateFood(foodId uuid.UUID, food *schemas.Food, ctx *gin.Context) error {
	if err := db.Model(&schemas.Food{}).Where("id = ?", foodId).Updates(food).Error; err != nil {
		logger.Errf("error updating food: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error updating food")
		return err
	}
	return nil
}

func DeleteFood(foodId uuid.UUID, ctx *gin.Context) error {
	if err := db.Delete(&schemas.Food{}, foodId).Error; err != nil {
		logger.Errf("error deleting food: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error deleting food")
		return err
	}
	return nil
}