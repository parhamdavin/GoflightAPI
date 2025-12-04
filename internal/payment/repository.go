package payment

import "example.com/airline-reservation/models"

type PaymentRepository interface {
	CraetePayment(payment *models.PaymentStatus)error
	GetPaymentByReservation(reservationID uint) ([]*models.PaymentStatus, error)
	GetPaymentByID(id uint)(*models.PaymentStatus,error)
	
}