package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yamess/inventory/controllers"
	"github.com/yamess/inventory/utils"
)

func ProductRoutes(r *gin.RouterGroup) {
	cat := r.Group("/product")
	{
		cat.POST("", utils.ProductValidator, controllers.CreateProduct)
		cat.GET("", controllers.GetProductList)
		cat.GET(":id", controllers.GetProductByID)
		cat.PATCH("/update/:id", controllers.UpdateProduct)
		cat.DELETE("/delete/:id", controllers.DeleteProduct)
	}
}
