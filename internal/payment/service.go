package payment

import "example.com/airline-reservation/models"

type PaymentService interface {
	CraetePayment(payment *models.PaymentStatus) error
	GetPaymentByReservation(reservationID uint) ([]*models.PaymentStatus, error)
	GetPaymentByID(id uint) (*models.PaymentStatus, error)
}

type paymentservice struct {
	repo PaymentRepository
}

// CraetePayment implements PaymentService.
func (p *paymentservice) CraetePayment(payment *models.PaymentStatus) error {
	return p.repo.CraetePayment(payment)
}

// GetPaymentByID implements PaymentService.
func (p *paymentservice) GetPaymentByID(id uint) (*models.PaymentStatus, error) {
	return p.repo.GetPaymentByID(id)
}

// GetPaymentByReservation implements PaymentService.
func (p *paymentservice) GetPaymentByReservation(reservationID uint) ([]*models.PaymentStatus, error) {
	return p.repo.GetPaymentByReservation(reservationID)
}

func NewPaymentService(repo PaymentRepository) PaymentService {
	return &paymentservice{repo: repo}
}
