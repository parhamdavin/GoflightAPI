package models

import "gorm.io/gorm"

// models/user.go
type User struct {
	gorm.Model
	Name         string
	SSID         string 
	PhoneNumber  string 
	PassportId   string 
	Email        string 
	Password     string
	Agency       bool
	Active       bool
	Role         string
	Reservations []Reservation `gorm:"foreignKey:UserID"`
	Wallet       Wallet
}
