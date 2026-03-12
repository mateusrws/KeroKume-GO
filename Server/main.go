package main

import (
	"kerokume-go/config"
	"kerokume-go/router"
)


var (
	logger *config.Logger
)

func main() {
	logger = config.GetLooger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errf("Config Initialization Err: %v", err)
		return
	}

	// Initialize Routes
	router.Initialize()
}
