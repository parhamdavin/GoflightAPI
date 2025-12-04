package models

import "time"

type ServiceOffering struct {
	FlightID  uint `gorm:"primaryKey"`
	ServiceID uint `gorm:"primaryKey"`
	OfferedYN bool
	FromDate  time.Time
	ToDate    time.Time
}
