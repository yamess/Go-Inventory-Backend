package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yamess/inventory/models"
	"log"
	"net/http"
)

//type KeyCategory struct{}

func Validator(c *gin.Context) {
	fmt.Println("I'm a dummy test test")
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		logString := fmt.Sprintf("Error while parsing the resquest json data.\n%s", err.Error())
		log.Println(logString)
		c.JSON(http.StatusBadRequest, "Error while parsing the request data")
		return
	}
	if err := category.Validate(); err != nil {
		logString := fmt.Sprintf("Invalid data schema.\n%s", err.Error())
		log.Println(logString)
		c.JSON(http.StatusBadRequest, "Invalid data schema")
		return
	}
	c.Set("category", category)
	c.Next()
}
