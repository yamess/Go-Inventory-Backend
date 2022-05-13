package models

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/yamess/inventory/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
	"log"
	"time"
)

type ProductAttributes struct {
	Id        uint   `json:"id" gorm:"primaryKey" example:"1"`
	ProductID uint   `json:"product_id" example:"1"`
	Name      string `json:"name" gorm:"primaryKey" example:"weight"`
	Value     string `json:"value" gorm:"primaryKey" example:"0.5"`
	Unit      string `json:"unit" gorm:"primaryKey" example:"kg"`
}
type ProductAttributesRequest struct {
	Name  string `json:"name" gorm:"primaryKey" example:"weight"`
	Value string `json:"value" gorm:"primaryKey" example:"0.5"`
	Unit  string `json:"unit" gorm:"primaryKey" example:"kg"`
}

type Product struct {
	Id                uint                `json:"id" gorm:"primaryKey" example:"1"`
	ProductName       string              `json:"product_name"  example:"TV Remote" validate:"required" gorm:"uniqueIndex:idx_member"`
	CategoryID        uint                `json:"category_id"  example:"1" validate:"required" gorm:"uniqueIndex:idx_member"`
	Description       string              `json:"description"  example:"Gen 2 tv remote" gorm:"uniqueIndex:idx_member"`
	CostPrice         float32             `json:"cost_price"  example:"12500" validate:"required" gorm:"uniqueIndex:idx_member"`
	SellingPrice      float32             `json:"selling_price"  example:"20000" gorm:"uniqueIndex:idx_member"`
	Quantity          int32               `json:"quantity"  example:"30" validate:"required"`
	SupplierID        uint                `json:"supplier_id"  example:"1"`
	ProductAttributes []ProductAttributes `json:"product_attributes" gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;cascade:all-delete-orphan;"`
	Base
}
type ProductRequest struct {
	ProductName       string                     `json:"product_name"  example:"TV Remote" validate:"required" gorm:"uniqueIndex:idx_member"`
	CategoryID        uint                       `json:"category_id"  example:"1" validate:"required" gorm:"uniqueIndex:idx_member"`
	Description       string                     `json:"description"  example:"Gen 2 tv remote" gorm:"uniqueIndex:idx_member"`
	CostPrice         float32                    `json:"cost_price"  example:"12500" validate:"required" gorm:"uniqueIndex:idx_member"`
	SellingPrice      float32                    `json:"selling_price"  example:"20000" gorm:"uniqueIndex:idx_member"`
	Quantity          int32                      `json:"quantity"  example:"30" validate:"required"`
	SupplierID        uint                       `json:"supplier_id"  example:"1"`
	ProductAttributes []ProductAttributesRequest `json:"product_attributes"`
}

type ProductList []Product

func (p *Product) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(&p)
}
func (p *ProductRequest) ToJSON() ([]byte, error) {
	data, err := json.Marshal(&p)
	return data, err
}

// Validate product schema
func (p *ProductRequest) Validate() error {
	validation := validator.New()
	return validation.Struct(p)
}

func (p *Product) CreateRecord(userId uint) *gorm.DB {
	p.CreatedBy = userId

	res := database.MyDB.Conn.
		Model(&p).
		Clauses(clause.Returning{}).
		Create(&p)
	return res
}
func (p *Product) GetRecord(searchField string, searchValue any) *gorm.DB {
	query := fmt.Sprintf("%s = ?", searchField)
	res := database.MyDB.Conn.Preload("ProductAttributes").Find(&p, query, searchValue).Limit(1)
	return res
}
func (p *ProductList) GetRecords() *gorm.DB {
	res := database.MyDB.Conn.Preload("ProductAttributes").Find(&p)
	return res
}
func (p *Product) UpdateRecord(id uint, userId uint) *gorm.DB {
	p.UpdatedBy = userId
	p.Id = id
	for _, v := range p.ProductAttributes {
		v.ProductID = id
	}
	err := database.MyDB.Conn.Model(&Product{Id: id}).Association("ProductAttributes").Clear()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Update product
	result := database.MyDB.Conn.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Clauses(clause.Returning{}, clause.OnConflict{UpdateAll: true}).
		Omit(
			"CreatedAt", "CreatedBy", "Id",
		).
		Updates(&p)

	// Delete orphan children
	res := database.MyDB.Conn.
		Where("product_id is ?", gorm.Expr("NULL")).
		Delete(&ProductAttributes{})
	if res.Error != nil {
		log.Println(res.Error)
	}

	return result
}
func (p *Product) DeleteRecord(id uint) *gorm.DB {
	p.Id = id
	res := database.MyDB.Conn.Select("ProductAttributes").Delete(&p)
	return res
}
func (p *ProductAttributes) AttributeBulkDelete() *gorm.DB {
	res := database.MyDB.Conn.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(ProductAttributes{})
	return res
}
func (p *Product) BulkDeleteRecord() *gorm.DB {
	res := database.MyDB.Conn.
		Session(&gorm.Session{AllowGlobalUpdate: true}).
		Select("ProductAttributes").
		Delete(Product{})
	return res
}

// BeforeCreate Hooks
func (p *Product) BeforeCreate(tx *gorm.DB) error {
	p.CreatedAt = time.Now().UTC()
	return nil
}
