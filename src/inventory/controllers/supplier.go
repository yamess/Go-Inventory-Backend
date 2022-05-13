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

// CreateSupplier ... Create Supplier
// @Summary Create new supplier based on parameters
// @Description Create new supplier
// @Tags Supplier
// @Accept json
// @Param Supplier body models.SupplierRequest true "Supplier Data"
// @Produce json
// @Success 200 {object} models.Supplier
// @Failure 400,404,500 {object} object
// @Router /supplier [post]
func CreateSupplier(c *gin.Context) {
	var supplier models.Supplier
	var supplierRequest models.SupplierRequest

	d, ok := c.Keys["supplierRequest"]
	if !ok {
		c.JSON(http.StatusBadRequest, "Enable to process request")
		return
	}
	supplierRequest = d.(models.SupplierRequest)
	data, err := supplierRequest.ToJSON()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, "Enable to process request")
		return
	}

	err = json.Unmarshal([]byte(data), &supplier)
	if err != nil {
		responseString := fmt.Sprintf("Supplier %s already exist in database", supplier.Name)
		log.Printf(responseString)
		c.JSON(http.StatusBadRequest, "Enable to process request")
		return
	}

	// Check if supplier already exist in db first
	res := supplier.GetRecord("name", supplier.Name)
	if res.RowsAffected > 0 {
		responseString := fmt.Sprintf("Supplier %s already exist in database", supplier.Name)
		log.Printf(responseString)
		c.JSON(http.StatusConflict, "supplier already exist")
		return
	}

	res = supplier.CreateRecord(configs.DefaultUser)
	if res.Error != nil {
		logString := fmt.Sprintf("Enable to save data into the database.\n%s", res.Error)
		log.Println(logString)
		c.JSON(http.StatusBadRequest, "Enable to save record")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &supplier, "length": 1})
}

// GetSupplierList ... Get list of suppliers
// @Summary Get all the suppliers
// @Description Get the list all the suppliers
// @Tags Supplier
// @Success 200 {array} models.Supplier
// @Failure 404 {object} object
// @Router /supplier [get]
func GetSupplierList(c *gin.Context) {
	var suppliers models.Suppliers

	res := suppliers.GetRecords()
	if res.Error != nil {
		logString := fmt.Sprintf("Enable to get data from database.\n%s", res.Error)
		log.Println(logString)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enable to get data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": suppliers, "length": len(suppliers)})
}

// GetSupplierByID ... Get the supplier by id
// @Summary Get one supplier
// @Description get supplier by ID
// @Tags Supplier
// @Param id path string true "Supplier ID"
// @Success 200 {object} models.Supplier
// @Failure 400,404 {object} object
// @Router /supplier/{id} [get]
func GetSupplierByID(c *gin.Context) {
	var supplier models.Supplier

	id, _ := strconv.Atoi(c.Param("id"))
	res := supplier.GetRecord("id", id)

	if res.Error != nil && !errors.Is(res.Error, gorm.ErrRecordNotFound) {
		logString := fmt.Sprintf("Enable to get data from database\n%s", res.Error)
		log.Println(logString)
		c.JSON(http.StatusBadRequest, gin.H{"error": res.Error, "message": "Enable to get data"})
		return
	} else if res.RowsAffected == 0 {
		logString := fmt.Sprintf("No record found for id: %d", id)
		log.Println(logString)
		c.JSON(http.StatusNotFound, gin.H{"error": "No record found", "message": "No record found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": supplier, "length": 1})
}

// UpdateSupplier ... Update Supplier
// @Summary Update existing supplier based on parameters
// @Description Update existing Supplier
// @Tags Supplier
// @Accept json
// @Param Supplier body models.Supplier true "Supplier Data"
// @Param id path string true "Supplier ID"
// @Produce json
// @Success 200 {object} models.Supplier
// @Failure 400,404,500 {object} object
// @Router /supplier/update/{id} [patch]
func UpdateSupplier(c *gin.Context) {
	var supplier *models.Supplier

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logString := fmt.Sprintf(err.Error())
		c.JSON(http.StatusBadRequest, logString)
		return
	}
	if err := c.BindJSON(&supplier); err != nil {
		logString := fmt.Sprintf("Enable to decode json data.\n%s", err.Error())
		log.Printf(logString)
		c.JSON(http.StatusBadRequest, "Enable to decode json data")
		return
	}
	res := supplier.UpdateRecord(uint(id), configs.DefaultUser)
	if res.Error != nil {
		logString := fmt.Sprintf("Enable to update record.\n%s", res.Error)
		log.Printf(logString)
		c.JSON(http.StatusBadRequest, "Enable to update record")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &supplier, "length": 1})
}

// DeleteSupplier ... Delete Supplier
// @Summary Delete existing supplier based on id
// @Description Delete existing supplier
// @Tags Supplier
// @Accept json
// @Param id path string true "Supplier ID"
// @Produce json
// @Success 200 {object} object
// @Failure 400,404,500 {object} object
// @Router /supplier/delete/{id} [delete]
func DeleteSupplier(c *gin.Context) {
	var supplier models.Supplier

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logString := fmt.Sprintf(err.Error())
		c.JSON(http.StatusBadRequest, logString)
		return
	}

	res := supplier.DeleteRecord(uint(id))
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

// BulkDeleteSupplier ... Delete all the Suppliers in the database
// @Summary Delete all the existing suppliers
// @Description Delete all the existing suppliers
// @Tags Supplier
// @Accept json
// @Produce json
// @Success 204 {object} object
// @Failure 400,404,500 {object} object
// @Router /supplier [delete]
func BulkDeleteSupplier(c *gin.Context) {
	var supplier models.Supplier
	res := supplier.BulkDeleteRecord()
	if res.Error != nil {
		logString := fmt.Sprintf("Enable to delete records from database.\n%s", res.Error)
		log.Printf(logString)
		c.JSON(http.StatusBadRequest, "Enable to delete records")
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
