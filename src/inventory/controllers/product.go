package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yamess/inventory/configs"
	"github.com/yamess/inventory/models"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

// CreateProduct ... Create Product
// @Summary Create new product based on parameters
// @Description Create new product
// @Tags Product
// @Accept json
// @Param Product body models.ProductRequest true "Product Data"
// @Produce json
// @Success 200 {object} models.Product
// @Failure 400,404,500 {object} object
// @Router /product [post]
func CreateProduct(c *gin.Context) {
	var product models.Product
	var productRequest models.ProductRequest

	d, ok := c.Keys["productRequest"]
	if !ok {
		c.JSON(http.StatusBadRequest, "Enable to process request")
		return
	}
	productRequest = d.(models.ProductRequest)
	data, err := productRequest.ToJSON()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, "Enable to process request")
		return
	}

	err = json.Unmarshal([]byte(data), &product)
	if err != nil {
		responseString := fmt.Sprintf("Produc %s already exist in database", product.ProductName)
		log.Printf(responseString)
		c.JSON(http.StatusBadRequest, "Enable to process request")
		return
	}

	// Check if supplier already exist in db first
	res := product.GetRecord("product_name", product.ProductName)
	if res.RowsAffected > 0 {
		responseString := fmt.Sprintf("Product %s already exist in database", product.ProductName)
		log.Printf(responseString)
		c.JSON(http.StatusConflict, "product already exist")
		return
	}

	res = product.CreateRecord(configs.DefaultUser)
	if res.Error != nil {
		logString := fmt.Sprintf("Enable to save data into the database.\n%s", res.Error)
		log.Println(logString)
		c.JSON(http.StatusBadRequest, "Enable to save record")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &product, "length": 1})
}

// GetProductList ... Get list of Products
// @Summary Get all the products
// @Description Get the list all the products
// @Tags Product
// @Success 200 {array} models.Product
// @Failure 404 {object} object
// @Router /product [get]
func GetProductList(c *gin.Context) {
	var productList models.ProductList

	res := productList.GetRecords()
	if res.Error != nil {
		logString := fmt.Sprintf("Enable to get data from database.\n%s", res.Error)
		log.Println(logString)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enable to get data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": productList, "length": len(productList)})
}

// GetProductByID ... Get the product by id
// @Summary Get one product
// @Description get product by ID
// @Tags Product
// @Param id path string true "Product ID"
// @Success 200 {object} models.Product
// @Failure 400,404 {object} object
// @Router /product/{id} [get]
func GetProductByID(c *gin.Context) {
	var product models.Product

	id, _ := strconv.Atoi(c.Param("id"))
	res := product.GetRecord("id", id)

	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		logString := fmt.Sprintf("Enable to get data from database\n%s", res.Error)
		log.Println(logString)
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error, "message": "Enable to get data"})
		return
	} else if res.RowsAffected == 0 {
		logString := fmt.Sprintf("No record found for id: %d", id)
		log.Println(logString)
		c.JSON(http.StatusNotFound, gin.H{"error": "No record found", "message": logString})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product, "length": 1})
}

// UpdateProduct ... Update Product
// @Summary Update existing product based on parameters
// @Description Update existing Product
// @Tags Product
// @Accept json
// @Param Product body models.Product true "Product Data"
// @Param id path string true "Product ID"
// @Produce json
// @Success 200 {object} models.Product
// @Failure 400,404,500 {object} object
// @Router /product/update/{id} [patch]
func UpdateProduct(c *gin.Context) {
	var product models.Product
	var tmpProduct models.Product

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logString := fmt.Sprintf(err.Error())
		c.JSON(http.StatusBadRequest, logString)
		return
	}

	// Check if record exist
	res := tmpProduct.GetRecord("id", id)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, "Could not get record from database")
		return
	} else if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, "Record not found")
		return
	}

	if err := c.BindJSON(&product); err != nil {
		logString := fmt.Sprintf("Enable to decode json data.\n%s", err.Error())
		log.Printf(logString)
		c.JSON(http.StatusBadRequest, "Enable to decode json data")
		return
	}

	res = product.UpdateRecord(uint(id), configs.DefaultUser)
	if res.Error != nil {
		logString := fmt.Sprintf("Enable to update record.\n%s", res.Error)
		log.Printf(logString)
		c.JSON(http.StatusBadRequest, "Enable to update record")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &product, "length": 1})
}

// DeleteProduct ... Delete Product
// @Summary Delete existing product based on id
// @Description Delete existing product
// @Tags Product
// @Accept json
// @Param id path string true "Product ID"
// @Produce json
// @Success 200 {object} object
// @Failure 400,404,500 {object} object
// @Router /product/delete/{id} [delete]
func DeleteProduct(c *gin.Context) {
	var product models.Product

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logString := fmt.Sprintf(err.Error())
		c.JSON(http.StatusBadRequest, logString)
		return
	}

	res := product.DeleteRecord(uint(id))
	if res.Error != nil {
		logString := fmt.Sprintf("Enable to delete record from database.\n%s", res.Error)
		log.Printf(logString)
		c.JSON(http.StatusBadRequest, "Enable to delete data")
		return
	} else if res.RowsAffected == 0 {
		logString := fmt.Sprint("No record found in database")
		log.Println(logString)
		c.JSON(http.StatusNotFound, "No record found")
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// BulkDeleteProduct ... Delete all the Products in the database
// @Summary Delete all the existing products
// @Description Delete all the existing products
// @Tags Product
// @Accept json
// @Produce json
// @Success 204 {object} object
// @Failure 400,404,500 {object} object
// @Router /product [delete]
func BulkDeleteProduct(c *gin.Context) {
	var product models.Product
	res := product.BulkDeleteRecord()
	if res.Error != nil {
		logString := fmt.Sprintf("Enable to delete records from database.\n%s", res.Error)
		log.Printf(logString)
		c.JSON(http.StatusBadRequest, "Enable to delete records")
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
