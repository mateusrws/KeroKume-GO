package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Food struct{
	gorm.Model
	name string
	description string
	price float32
	pathImg string
	foodCategory string
	isAvailable bool
	menuId uuid.UUID
}