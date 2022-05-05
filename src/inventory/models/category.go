package models

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/yamess/inventory/database"
	"gorm.io/gorm"
	"io"
	"time"
)

type Category struct {
	Id   uint   `json:"id" gorm:"primaryKey" example:"1"`
	Name string `json:"name" gorm:"unique" validate:"required" example:"Phone"`
	Base
}

type Categories []Category

// Validate Schema Validation
func (c *Category) Validate() error {
	validation := validator.New()
	return validation.Struct(c)
}

func (c *Category) FromJSON(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&c)
	return err
}

func (c *Category) CreateRecord(userId uint) *gorm.DB {
	c.CreatedBy = userId
	res := database.MyDB.Conn.Create(&c)
	return res
}
func (c *Category) GetRecord(searchField string, searchValue any) *gorm.DB {
	query := fmt.Sprintf("%s = ?", searchField)
	result := database.MyDB.Conn.Limit(1).Find(&c, query, searchValue)
	return result
}
func (cs *Categories) GetRecords() *gorm.DB {
	res := database.MyDB.Conn.Find(&cs)
	return res
}
func (c *Category) UpdateRecord(id uint, userId uint) *gorm.DB {
	c.UpdatedAt.Time = time.Now().UTC()
	c.UpdatedBy = userId
	c.Id = id
	res := database.MyDB.Conn.Model(&c).Omit("id", "created_at", "created_by").Updates(c)
	return res
}
func (c *Category) DeleteRecord(id uint) *gorm.DB {
	c.Id = id
	res := database.MyDB.Conn.Delete(&c)
	return res
}

// BeforeCreate Hooks
func (c *Category) BeforeCreate(tx *gorm.DB) error {
	c.CreatedAt = time.Now().UTC()
	return nil
}
