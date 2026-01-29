package schemas

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Menu struct {
	gorm.Model
	id					 uuid.UUID
	name         string
	restaurantId uuid.UUID
}
