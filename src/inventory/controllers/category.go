package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yamess/inventory/models"
	"log"
	"net/http"
)

func GetCategoryList(c *gin.Context) {
	var categories models.Categories

	res := categories.GetRecords()
	if res.Error != nil {
		logString := fmt.Sprintf("Enable to get data from database.\n%s", res.Error)
		log.Println(logString)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enable to get data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}
