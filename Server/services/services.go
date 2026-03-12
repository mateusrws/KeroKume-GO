package services

import "kerokume-go/config"

var(
	logger *config.Logger
)

func Init(){
	logger = config.GetLooger("handler")
}