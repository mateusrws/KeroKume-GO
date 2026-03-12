package repos

import (
	"gorm.io/gorm"

	"kerokume-go/config"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init(){
	logger = config.GetLooger("handler")
	db = config.GetDB()
}
