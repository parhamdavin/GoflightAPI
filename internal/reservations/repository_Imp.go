package reservations

import (
	"example.com/airline-reservation/models"
	"gorm.io/gorm"
)

type reservationrepository struct {
	db *gorm.DB
}

// CreateReservation implements ReservationRepository.
func (r *reservationrepository) CreateReservation(reservation *models.Reservation) error {
	if err := r.db.Create(reservation).Error; err != nil {
		return err
	}
	return nil
}

// GetReservationByID implements ReservationRepository.
func (r *reservationrepository) GetReservationByID(id uint) (*models.Reservation, error) {
	var reservation models.Reservation

	if err := r.db.First(&reservation, id).Error; err != nil {
		return nil, err
	}

	return &reservation, nil
}

// GetReservationsByUser implements ReservationRepository.
func (r *reservationrepository) GetReservationsByUser(userID uint) ([]*models.Reservation, error) {
	var reservations []*models.Reservation

	if err :=r.db.Where("UserID = ?", userID).Find(&reservations).Error; err != nil {
		return nil, err
	}

	return reservations, nil
}

func NewReservationRepository(db *gorm.DB) ReservationRepository {
	return &reservationrepository{db}
}
