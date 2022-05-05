package main

import (
	"github.com/yamess/inventory/app"
	"github.com/yamess/inventory/configs"
	"github.com/yamess/inventory/database"
)

func main() {
	// Load environment variables
	configs.InitEnv()

	// Apply auto migration of the models
	database.Automigrate()

	app.Run()
}
