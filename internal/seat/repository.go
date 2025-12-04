package seat

import "example.com/airline-reservation/models"

type SeatRepository interface {
	GetSeatsByFlightID(flightID uint) ([]*models.SeatDetail, error)
	ReserveSeat(seatID uint, reservationID uint) error
	ReleaseSeat(seatID uint) error
	GetSeatByID(seatID uint) (*models.SeatDetail, error)
	UpdateSeat(seat *models.SeatDetail) error
	CheckIfSeatAvailable(seatID uint) (bool, error)
}