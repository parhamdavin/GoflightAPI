package reservations

import "example.com/airline-reservation/models"

type ReservationService interface {
	CreateReservation(reservation *models.Reservation) error
	GetReservationByID(id uint) (*models.Reservation, error)
	GetReservationsByUser(userID uint) ([]*models.Reservation, error)
}

type reservationservice struct {
	repo ReservationRepository
}

// CreateReservation implements ReservationService.
func (r *reservationservice) CreateReservation(reservation *models.Reservation) error {
	return r.repo.CreateReservation(reservation)
}

// GetReservationByID implements ReservationService.
func (r *reservationservice) GetReservationByID(id uint) (*models.Reservation, error) {
	return r.repo.GetReservationByID(id)
}

// GetReservationsByUser implements ReservationService.
func (r *reservationservice) GetReservationsByUser(userID uint) ([]*models.Reservation, error) {
	return r.repo.GetReservationsByUser(userID)
}

func NewReservationService(repo ReservationRepository) ReservationService {
	return &reservationservice{repo: repo}
}
