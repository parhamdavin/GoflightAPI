package models

import (
	"time"
	"gorm.io/gorm"
)

type PaymentStatus struct {
	gorm.Model	
	
	
	ReservationID uint `gorm:"not null"`// Reservation ID

	
	Amount float64 `gorm:"type:decimal(10,2);not null"`// Payment amount

	
	PaymentStatus string `gorm:"type:varchar(50);not null"`// Payment status (e.g. 'Pending', 'Completed')

	
	PaymentMethod string `gorm:"type:varchar(50);not null"`// Payment method 

	
	PaymentDate time.Time `gorm:"type:timestamp;not null"`// Payment date and time
	Reservation   Reservation `gorm:"foreignKey:ReservationID"`
}
