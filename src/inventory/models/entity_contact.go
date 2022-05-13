package models

type Contact struct {
	Id       uint   `json:"id" gorm:"primaryKey" example:"1"`
	EntityID uint   `json:"entity_id" gorm:"unique" example:"1"`
	Phone    string `json:"phone" validate:"required" example:"+22675010203"`
	Email    string `json:"email" validate:"email"  example:"fake@faker.com"`
	Fax      string `json:"fax" example:"+22675010203"`
}

type ContactRequest struct {
	Phone string `json:"phone" validate:"required" example:"+22675010203"`
	Email string `json:"email" validate:"email"  example:"fake@faker.com"`
	Fax   string `json:"fax" example:"+22675010203"`
}
