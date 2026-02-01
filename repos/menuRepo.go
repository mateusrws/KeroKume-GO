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
