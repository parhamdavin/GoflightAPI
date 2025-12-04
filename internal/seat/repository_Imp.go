package seat

import (
	"fmt"

	"example.com/airline-reservation/models"
	"gorm.io/gorm"
)

type seatrepository struct {
	db *gorm.DB
}

// CheckIfSeatAvailable implements SeatRepository.
func (s *seatrepository) CheckIfSeatAvailable(seatID uint) (bool, error) {
	var seat *models.SeatDetail
	if err := s.db.First(&seat, seatID).Error; err != nil {
		return false, err
	}
	if seat.ReservationID == 0 {
		return true, nil 
	}
	return false, nil
}

// GetSeatByID implements SeatRepository.
func (s *seatrepository) GetSeatByID(seatID uint) (*models.SeatDetail, error) {
	var seat models.SeatDetail
	if err := s.db.First(&seat, seatID).Error; err != nil {
		return nil, err
	}
	return &seat, nil
}


// GetSeatsByFlightID implements SeatRepository.
func (s *seatrepository) GetSeatsByFlightID(flightID uint) ([]*models.SeatDetail, error) {
	var seats []*models.SeatDetail
	if err := s.db.Where("FlightID = ?", flightID).Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}


// ReleaseSeat implements SeatRepository.
func (s *seatrepository) ReleaseSeat(seatID uint) error {
	var seat models.SeatDetail
	if err := s.db.First(&seat, seatID).Error; err != nil {
		return err
	}
	seat.ReservationID = 0
	if err := s.db.Save(&seat).Error; err != nil {
		return err
	}
	return nil
}


// ReserveSeat implements SeatRepository.
func (s *seatrepository) ReserveSeat(seatID uint, reservationID uint) error {
	var seat models.SeatDetail
	if err := s.db.First(&seat, seatID).Error; err != nil {
		return err
	}
	if seat.ReservationID != 0 {
		return fmt.Errorf("this seat is already reserved")
	}
	seat.ReservationID = reservationID
	if err := s.db.Save(&seat).Error; err != nil {
		return err
	}
	return nil
}


// UpdateSeat implements SeatRepository.
func (s *seatrepository) UpdateSeat(seat *models.SeatDetail) error {
	if err := s.db.Save(seat).Error; err != nil {
		return err
	}
	return nil
}


func NewSeatRepository(db *gorm.DB) SeatRepository {
	return &seatrepository{db}
}
