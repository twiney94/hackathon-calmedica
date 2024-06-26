package models

type User struct {
	ID          string `gorm:"primaryKey"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
}
