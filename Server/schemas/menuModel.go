package schemas

import (
	"github.com/google/uuid"
)

type Menu struct {
	BaseModel
	Name         string
	RestaurantId uuid.UUID
}
