package controllers

import (
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

// CreateCategory ... Create Category
// @Summary Create new category based on parameters
// @Description Create new category
// @Tags Category
// @Accept json
// @Param Category body models.CategoryRequest true "Category Data"
// @Produce json
// @Success 200 {object} models.Category
// @Failure 400,404,500 {object} object
// @Router /category [post]
func CreateCategory(c *gin.Context) {
	var category models.Category
	category = c.Keys["category"].(models.Category)

	// Check if category already exist in db first
	res := category.GetRecord("name", category.Name)
	if res.RowsAffected > 0 {
		responseString := fmt.Sprintf("Category %s already exist in database", category.Name)
		log.Printf(responseString)
		c.JSON(http.StatusConflict, "Category already exist")
		return
	}

	res = category.CreateRecord(configs.DefaultUser)
	if res.Error != nil {
		logString := fmt.Sprintf("Enable to save data into the database.\n%s", res.Error)
		log.Println(logString)
		c.JSON(http.StatusBadRequest, "Enable to save record")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": category, "length": 1})
}

// GetCategoryList ... Get list of categories
// @Summary Get all the categories
// @Description Get the list all the categories
// @Tags Category
// @Success 200 {array} models.Category
// @Failure 404 {object} object
// @Router /category [get]
func GetCategoryList(c *gin.Context) {
	var categories models.Categories

	res := categories.GetRecords()
	if res.Error != nil {
		logString := fmt.Sprintf("Enable to get data from database.\n%s", res.Error)
		log.Println(logString)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Enable to get data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories, "length": len(categories)})
}

// GetCategoryByID ... Get the category by id
// @Summary Get one category
// @Description get category by ID
// @Tags Category
// @Param id path string true "Category ID"
// @Success 200 {object} models.Category
// @Failure 400,404 {object} object
// @Router /category/{id} [get]
func GetCategoryByID(c *gin.Context) {
	var category models.Category

	id, _ := strconv.Atoi(c.Param("id"))
	res := category.GetRecord("id", id)

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
	c.JSON(http.StatusOK, gin.H{"data": category, "length": 1})
}

// UpdateCategory ... Update Category
// @Summary Update existing category based on parameters
// @Description Update existing category
// @Tags Category
// @Accept json
// @Param Category body models.CategoryRequest true "Category Data"
// @Param id path string true "Category ID"
// @Produce json
// @Success 200 {object} models.Category
// @Failure 400,404,500 {object} object
// @Router /category/update/{id} [patch]
func UpdateCategory(c *gin.Context) {
	var category *models.Category

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logString := fmt.Sprintf(err.Error())
		c.JSON(http.StatusBadRequest, logString)
		return
	}
	if err := c.BindJSON(&category); err != nil {
		logString := fmt.Sprintf("Enable to decode json data.\n%s", err.Error())
		log.Printf(logString)
		c.JSON(http.StatusBadRequest, "Enable to decode json data")
		return
	}

	res := category.UpdateRecord(uint(id), configs.DefaultUser)
	if res.Error != nil {
		logString := fmt.Sprintf("Enable to update record.\n%s", res.Error)
		log.Printf(logString)
		c.JSON(http.StatusBadRequest, "Enable to update record")
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": &category, "length": 1})
}

// DeleteCategory ... Delete Category
// @Summary Delete existing category based on id
// @Description Delete existing category
// @Tags Category
// @Accept json
// @Param id path string true "Category ID"
// @Produce json
// @Success 200 {object} object
// @Failure 400,404,500 {object} object
// @Router /category/delete/{id} [delete]
func DeleteCategory(c *gin.Context) {
	var category models.Category

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logString := fmt.Sprintf(err.Error())
		c.JSON(http.StatusBadRequest, logString)
		return
	}

	res := category.DeleteRecord(uint(id))
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
