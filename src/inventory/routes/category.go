package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yamess/inventory/controllers"
	"github.com/yamess/inventory/utils"
)

func CategoryRoutes(r *gin.RouterGroup) {
	cat := r.Group("/category")
	{
		cat.POST("", utils.CategoryValidator, controllers.CreateCategory)
		cat.GET("", controllers.GetCategoryList)
		cat.GET(":id", controllers.GetCategoryByID)
		cat.PATCH("/update/:id", controllers.UpdateCategory)
		cat.DELETE("/delete/:id", controllers.DeleteCategory)
	}
}
