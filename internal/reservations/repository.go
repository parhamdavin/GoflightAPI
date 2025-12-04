package reservations

import "example.com/airline-reservation/models"

type ReservationRepository interface {
	CreateReservation(reservation *models.Reservation) error
	GetReservationByID(id uint) (*models.Reservation, error)
	GetReservationsByUser(userID uint) ([]*models.Reservation, error)
	
}