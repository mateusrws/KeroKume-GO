package schemas

import (
	"gorm.io/gorm"
)

type Restaurant struct{
	gorm.Model
	name string
	password string
	description string
	
}