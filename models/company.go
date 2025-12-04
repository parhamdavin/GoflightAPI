package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string 
	LogoURL   string 
	Country   string
	Flights []FlightDetail
	 
}