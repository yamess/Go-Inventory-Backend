package main

import (
	"github.com/yamess/inventory/app"
	"github.com/yamess/inventory/configs"
	"github.com/yamess/inventory/database"
	"github.com/yamess/inventory/models"
)

func main() {
	// @title User API documentation
	// @version 1.0.0
	// @description     This is a simple rest api for category
	// @contact.name   API Support
	// @contact.url    http://www.swagger.io/support
	// @contact.email  support@swagger.io
	// @host localhost:8081
	// @BasePath /api/v1

	// Load environment variables
	configs.InitEnv()

	// Apply auto migration of the models
	database.Automigrate(
		models.Category{},
		models.Supplier{}, models.Address{}, models.Contact{},
		models.ProductAttributes{}, models.Product{},
	)

	app.Run()
}
