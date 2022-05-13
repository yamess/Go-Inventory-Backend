package models

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/yamess/inventory/database"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io"
	"time"
)

type Supplier struct {
	Id              uint    `json:"id" gorm:"primaryKey" example:"1"`
	Name            string  `json:"name" validate:"required" example:"Analytica Inc."`
	PersonToContact string  `json:"person_to_contact" example:"Willy Fatime"`
	Contact         Contact `json:"contact" gorm:"foreignKey:EntityID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" validate:"required"`
	Address         Address `json:"address" gorm:"foreignKey:EntityID; constraint:OnUpdate:CASCADE,OnDelete:CASCADE" validate:"required"`
	Base
}
type SupplierRequest struct {
	Name            string         `json:"name" validate:"required" example:"Analytica Inc."`
	PersonToContact string         `json:"person_to_contact" example:"Willy Fatime"`
	Contact         ContactRequest `json:"contact" validate:"required"`
	Address         AddressRequest `json:"address" validate:"required"`
}

type Suppliers []Supplier

// FromJSON to create struct
func (s *Supplier) FromJSON(r io.Reader) error {
	err := json.NewDecoder(r).Decode(&s)
	return err
}
func (s *SupplierRequest) ToJSON() ([]byte, error) {
	data, err := json.Marshal(&s)
	return data, err
}

// Validate supplier schema
func (s *SupplierRequest) Validate() error {
	validation := validator.New()
	return validation.Struct(s)
}

func (s *Supplier) CreateRecord(userId uint) *gorm.DB {
	s.CreatedBy = userId

	res := database.MyDB.Conn.
		Model(&s).
		Clauses(clause.Returning{}).
		Create(&s)
	return res
}
func (s *Supplier) GetRecord(searchField string, searchValue any) *gorm.DB {
	query := fmt.Sprintf("%s = ?", searchField)
	res := database.MyDB.Conn.Preload("Address").Preload("Contact").Find(&s, query, searchValue)
	return res
}
func (s *Suppliers) GetRecords() *gorm.DB {
	res := database.MyDB.Conn.Preload("Address").Preload("Contact").Find(&s)
	return res
}
func (s *Supplier) UpdateRecord(id uint, userId uint) *gorm.DB {
	s.UpdatedBy = userId
	s.Id = id
	res := database.MyDB.Conn.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Clauses(clause.Returning{}).
		Omit(
			"Id", "CreatedAt", "CreatedBy", "Contact.Id", "Contact.EntityID",
			"Address.Id", "Address.EntityID",
		).
		Updates(&s)
	return res
}
func (s *Supplier) DeleteRecord(id uint) *gorm.DB {
	s.Id = id
	res := database.MyDB.Conn.Delete(&s)
	return res
}
func (s *Supplier) BulkDeleteRecord() *gorm.DB {
	res := database.MyDB.Conn.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(Supplier{})
	return res
}

// BeforeCreate Hooks
func (s *Supplier) BeforeCreate(tx *gorm.DB) error {
	s.CreatedAt = time.Now().UTC()
	return nil
}
