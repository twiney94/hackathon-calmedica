package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID `gorm:"unique;type:uuid;primaryKey"`
	FirstName   string    `gorm:"not null" json:"firstName"`
	LastName    string    `gorm:"not null" json:"lastName"`
	PhoneNumber string    `gorm:"not null" json:"phoneNumber"`
}
