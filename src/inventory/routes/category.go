package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yamess/inventory/controllers"
)

func CategoryRoutes(r *gin.RouterGroup) {
	cat := r.Group("/category")
	{
		cat.GET("", controllers.GetCategoryList)
	}
}
