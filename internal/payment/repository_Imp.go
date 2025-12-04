package payment

import (
	"example.com/airline-reservation/models"
	"gorm.io/gorm"
)

type paymentrepository struct {
	db *gorm.DB
}

// GetPaymentByID implements PaymentRepository.
func (p *paymentrepository) GetPaymentByID(id uint) (*models.PaymentStatus, error) {
	var payment models.PaymentStatus

	if err := p.db.First(&payment, id).Error; err != nil {
		return nil, err
	}

	return &payment, nil
}


func (p *paymentrepository) CraetePayment(payment *models.PaymentStatus) error {
	if err := p.db.Create(payment).Error; err != nil {
		return err
	}
	return nil
}

// GetPaymentByReservation implements PaymentRepository.
func (p *paymentrepository) GetPaymentByReservation(reservationID uint) ([]*models.PaymentStatus, error) {
	var payments []*models.PaymentStatus

	if err := p.db.Where("ReservationID = ?", reservationID).Find(&payments).Error; err != nil {
		return nil, err
	}

	return payments, nil
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentrepository{db}
}
