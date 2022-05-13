package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yamess/inventory/controllers"
	"github.com/yamess/inventory/utils"
)

func SupplierRoutes(r *gin.RouterGroup) {
	supplier := r.Group("/supplier")
	{
		supplier.POST("", utils.SupplierValidator, controllers.CreateSupplier)
		supplier.GET("", controllers.GetSupplierList)
		supplier.DELETE("", controllers.BulkDeleteSupplier)
		supplier.GET(":id", controllers.GetSupplierByID)
		supplier.PUT("/update/:id", controllers.UpdateSupplier)
		supplier.DELETE("/delete/:id", controllers.DeleteSupplier)
	}
}
