package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yamess/inventory/models"
	"log"
	"net/http"
)

func CategoryValidator(c *gin.Context) {
	var category models.Category

	if err := c.BindJSON(&category); err != nil {
		logString := fmt.Sprintf("Error while parsing the resquest json data.\n%s", err)
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

func SupplierValidator(c *gin.Context) {
	var supplierRequest models.SupplierRequest

	if err := c.BindJSON(&supplierRequest); err != nil {
		logString := fmt.Sprintf("Error while parsing the resquest json data.\n%s", err)
		log.Println(logString)
		c.JSON(http.StatusBadRequest, "Error while parsing the request data")
		return
	}
	if err := supplierRequest.Validate(); err != nil {
		logString := fmt.Sprintf("Invalid data schema.\n%s", err.Error())
		log.Println(logString)
		c.JSON(http.StatusBadRequest, "Invalid data schema")
		return
	}

	c.Set("supplierRequest", supplierRequest)
	c.Next()
}

func ProductValidator(c *gin.Context) {
	var productRequest models.ProductRequest

	if err := c.BindJSON(&productRequest); err != nil {
		logString := fmt.Sprintf("Error while parsing the resquest json data.\n%s", err)
		log.Println(logString)
		c.JSON(http.StatusBadRequest, "Error while parsing the request data")
		return
	}
	if err := productRequest.Validate(); err != nil {
		logString := fmt.Sprintf("Invalid data schema.\n%s", err.Error())
		log.Println(logString)
		c.JSON(http.StatusBadRequest, "Invalid data schema")
		return
	}

	c.Set("productRequest", productRequest)
	c.Next()
}

func SetupCORS(ctx *gin.Context) {
	ctx.Request.Header.Set("Content-Type", "application/json")
	ctx.Request.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Request.Header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length,"+
		"Accept-Encoding, X-CSRF-Token, Authorization")
	ctx.Next()
}
