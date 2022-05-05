package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yamess/inventory/configs"
	"github.com/yamess/inventory/routes"
	"github.com/yamess/inventory/utils"
	"log"
)

func Run() {
	baseRoute := gin.Default()

	// Defining middlewares before any other code
	baseRoute.Use(utils.Validator)

	basePath := fmt.Sprintf("/api/v%s", configs.Version)
	v1 := baseRoute.Group(basePath)

	// Apply Category routers
	routes.CategoryRoutes(v1)

	// Run the server
	err := baseRoute.Run(configs.Host)
	if err != nil {
		log.Panic(err.Error())
	}
}
