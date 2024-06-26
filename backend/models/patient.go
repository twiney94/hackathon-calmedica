package models

import "time"

type Patient struct {
	UUID      string     `json:"uuid" gorm:"type:char(36);primary_key;index;unique"`
	Status    string     `json:"status" gorm:"default:bleu"`
	Phone     string     `json:"phone" gorm:"size:20"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" sql:"index"`
}

type PatientDTO struct {
	UUID      string     `json:"uuid"`
	Status    string     `json:"status"`
	Phone     string     `json:"phone"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}