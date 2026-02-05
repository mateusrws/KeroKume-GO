package repos

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"kerokume-go/schemas"
	"kerokume-go/utils"
)

func SaveMenu(menu *schemas.Menu) error {
	if err := db.Create(menu).Error; err != nil {
		logger.Errf("error creating menu: %v", err)
		return err
	}
	return nil
}
func FindUniqueMenu(id uuid.UUID, ctx *gin.Context) (schemas.Menu, error) {
	var menu schemas.Menu
	if err := db.First(&menu, id).Error; err != nil {
		logger.Errf("error finding menu: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error finding menu on database")
		return schemas.Menu{}, err
	}
	return menu, nil
}

func FindAllMenu(ctx *gin.Context) ([]schemas.Menu, error) {
	var menus []schemas.Menu
	if err := db.Find(&menus).Error; err != nil {
		logger.Errf("error finding all menus: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error finding all menus")
		return nil, err
	}
	return menus, nil
}

func FindAllByRestaurantId(restaurantId uuid.UUID, ctx *gin.Context) ([]schemas.Menu, error) {
	var menus []schemas.Menu
	if err := db.Where("restaurant_id = ?", restaurantId).Find(&menus).Error; err != nil {
		logger.Errf("error finding all menus by restaurant id: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error finding all menus by restaurant id")
		return nil, err
	}
	return menus, nil
}

func UpdateMenu(menuId uuid.UUID, menu *schemas.Menu, ctx *gin.Context) error {
	if err := db.Model(&schemas.Menu{}).Where("id = ?", menuId).Updates(menu).Error; err != nil {
		logger.Errf("error updating menu: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error updating menu")
		return err
	}
	return nil
}

func DeleteMenu(menuId uuid.UUID, ctx *gin.Context) error {
	if err := db.Delete(&schemas.Menu{}, menuId).Error; err != nil {
		logger.Errf("error deleting menu: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error deleting menu")
		return err
	}
	return nil
}