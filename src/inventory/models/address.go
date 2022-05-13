package models

type Address struct {
	Id           uint   `json:"id" gorm:"primaryKey" example:"1"`
	EntityID     uint   `json:"entity_id" gorm:"unique" example:"1"`
	AddressLine1 string `json:"address_line_1" example:"25 Rue Charles De Gaulle"`
	AddressLine2 string `json:"address_line_2" example:"Apartment 103"`
	City         string `json:"ville" example:"Ouagadougou" validate:"required"`
	Region       string `json:"region" example:"Kadiogo"`
	Country      string `json:"country" example:"Burkina Faso" validate:"required"`
	PostalCode   string `json:"postal_code" example:"1 BP 1023 Ouagadougou 1"`
}

type AddressRequest struct {
	AddressLine1 string `json:"address_line_1" example:"25 Rue Charles De Gaulle"`
	AddressLine2 string `json:"address_line_2" example:"Apartment 103"`
	City         string `json:"ville" example:"Ouagadougou" validate:"required"`
	Region       string `json:"region" example:"Kadiogo"`
	Country      string `json:"country" example:"Burkina Faso" validate:"required"`
	PostalCode   string `json:"postal_code" example:"1 BP 1023 Ouagadougou 1"`
}
