package models

import (
	"time"
)

type Notification struct {
	ID        uint   `gorm:"primaryKey"`
	PatientID string `gorm:"type:char(36);index"`
	Message   string `gorm:"type:text"`
	CreatedAt time.Time
}
