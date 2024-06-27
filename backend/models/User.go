package models

type User struct {
	ID          string `gorm:"unique;primaryKey"`
	FirstName   string `gorm:"not null" json:"firstName"`
	LastName    string `gorm:"not null" json:"lastName"`
	PhoneNumber string `gorm:"not null" json:"phoneNumber"`
}
