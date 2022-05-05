package models

import "time"

type Base struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt NullTime  `json:"updated_at" gorm:"default:null"`
	CreatedBy uint      `json:"created_by" example:"1"`
	UpdatedBy uint      `json:"updated_by" gorm:"default:null"`
}
