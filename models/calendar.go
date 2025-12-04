package models

import "time"

type Calendar struct {
	ID      uint `gorm:"primaryKey"`
	Date    time.Time
	Flights []FlightDetail `gorm:"foreignKey:CalendarID"`
}