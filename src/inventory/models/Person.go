package models

import "github.com/gofrs/uuid"

type Person struct {
	ID          uint
	UserID      uuid.UUID
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string

	Base
}
