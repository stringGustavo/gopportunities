package main

import (
	"github.com/stringGustavo/gopportunities.git/config"
	"github.com/stringGustavo/gopportunities.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	if err := config.Init(); err != nil {
		logger.Errorf("config initialization error %v",err)
		return
	}

	router.Initialize()
}
