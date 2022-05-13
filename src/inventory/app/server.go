package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/yamess/inventory/configs"
	_ "github.com/yamess/inventory/docs"
	"github.com/yamess/inventory/routes"
	"github.com/yamess/inventory/utils"
	"log"
)

func Run() {
	// This is for release in production mode
	//gin.SetMode(gin.ReleaseMode)
	baseRoute := gin.Default()
	baseRoute.Use(utils.SetupCORS)
	// Defining middlewares before any other code
	//baseRoute.Use(utils.Validator)

	basePath := fmt.Sprintf("/api/v%s", configs.Version)
	v1 := baseRoute.Group(basePath)

	// Apply Category routers
	routes.CategoryRoutes(v1)
	// Apply Supplier routes
	routes.SupplierRoutes(v1)
	// Apply product Routes
	routes.ProductRoutes(v1)

	// Setting Swagger doc endpoint
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the server
	err := baseRoute.Run(configs.Host)
	if err != nil {
		log.Panic(err.Error())
	}
}
