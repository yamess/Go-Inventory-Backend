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

	c.JSON(http.StatusOK, gin.H{"data": categories})
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
	c.JSON(http.StatusOK, gin.H{"data": category})
}

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
	user := configs.DefaultUser

	var category models.Category
	category := c.(*model)
}
