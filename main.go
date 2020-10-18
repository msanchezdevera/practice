package main

import (
	"practice/pkg/app"
	"practice/pkg/config"
	"practice/pkg/log"
)

func main() {

	config := config.ParseConfig()

	logger := log.NewLogger(config)

	logger.Infof("Launching practice")

	application := app.NewApplication(config, logger)

	application.Start()

}
