package models

import "gorm.io/gorm"

type Airport struct {
	gorm.Model
	Name           string
	AirportCity    string
	AirportCountry string
	FlightsFrom    []FlightDetail `gorm:"foreignKey:SourceID"`
	FlightsTo      []FlightDetail `gorm:"foreignKey:DestinationID"`
}
