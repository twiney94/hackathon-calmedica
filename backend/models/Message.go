package models

import (
	"time"
)

type Message struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	PatientId int       `json:"patient_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
