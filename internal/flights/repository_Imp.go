package flights

import (
	"fmt"

	"example.com/airline-reservation/models"
	"gorm.io/gorm"
)

type flightRepository struct {
	db *gorm.DB
}

// CreateCalendar implements FlightRespository.
func (f *flightRepository) CreateCalendar(calendar *models.Calendar) error {
	err := f.db.Create(calendar)
	if err != nil {
		return err.Error
	}
	return nil
}

// DeleteCalendar implements FlightRespository.
func (f *flightRepository) DeleteCalendar(id uint) error {
	result := f.db.Delete(&models.Calendar{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllCalendar implements FlightRespository.
func (f *flightRepository) GetAllCalendar() ([]*models.Calendar, error) {
	var calendar []*models.Calendar
	if err := f.db.Find(&calendar).Error; err != nil {
		return nil, err
	}
	return calendar, nil
}

// GetCalendarByID implements FlightRespository.
func (f *flightRepository) GetCalendarByID(id uint) (*models.Calendar, error) {
	calendar := &models.Calendar{}

	if err := f.db.First(calendar, id).Error; err != nil {
		return nil, err
	}

	return calendar, nil
}

// UpdateCalendar implements FlightRespository.
func (f *flightRepository) UpdateCalendar(calendar *models.Calendar) error {
	if err := f.db.Save(calendar).Error; err != nil {
		return err
	}
	return nil
}

// CreateFlightCost implements FlightRespository.
func (f *flightRepository) CreateFlightCost(cost *models.FlightCost) error {
	
	err := f.db.Create(cost)
	if err != nil {
		return err.Error
	}
	return nil
}

// DeleteFlightCost implements FlightRespository.
func (f *flightRepository) DeleteFlightCost(id uint) error {
	result := f.db.Delete(&models.FlightCost{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllFlightCost implements FlightRespository.
func (f *flightRepository) GetAllFlightCost() ([]*models.FlightCost, error) {
	var cost []*models.FlightCost
	if err := f.db.Find(&cost).Error; err != nil {
		return nil, err
	}
	return cost, nil
}

// GetFlightCostServiceByID implements FlightRespository.
func (f *flightRepository) GetFlightCostByID(id uint) (*models.FlightCost, error) {
	cost := &models.FlightCost{}

	if err := f.db.First(cost, id).Error; err != nil {
		return nil, err
	}

	return cost, nil
}

// UpdateFlightCost implements FlightRespository.
func (f *flightRepository) UpdateFlightCost(cost *models.FlightCost) error {
	if err := f.db.Save(cost).Error; err != nil {
		return err
	}
	return nil
}

func (f *flightRepository) GetFlightByID(id uint) (*models.FlightDetail, error) {
	flight := &models.FlightDetail{}

	if err := f.db.First(flight, id).Error; err != nil {
		return nil, err
	}

	return flight, nil
}

func (f *flightRepository) GetFlightServiceByID(id uint) (*models.FlightService, error) {
	flightservice := &models.FlightService{}

	if err := f.db.First(flightservice, id).Error; err != nil {
		return nil, err
	}

	return flightservice, nil
}

func NewFlightRepository(db *gorm.DB) FlightRespository {
	return &flightRepository{db}
}

func (f *flightRepository) CreateFlight(flight *models.FlightDetail) error {
	// مرحله 1: ایجاد پرواز
	if err := f.db.Create(&flight).Error; err != nil {
		return err
	}

	// مرحله 2: ایجاد صندلی‌ها
	var seats []models.SeatDetail
	for i := 1; i <= int(flight.SeatCount); i++ {
		seats = append(seats, models.SeatDetail{
			SeatNumber: fmt.Sprintf("S-%03d", i), // مثل S-001, S-002,...
			FlightID:   flight.ID,
			IsReserved: false, // فرض می‌کنیم که صندلی‌ها در ابتدا رزرو نشده‌اند
		})
	}

	// مرحله 3: ذخیره صندلی‌ها در دیتابیس
	if err := f.db.Create(&seats).Error; err != nil {
		return err
	}

	return nil
}


func (f *flightRepository) GetAllFlight() ([]*models.FlightDetail, error) {
	var flights []*models.FlightDetail
	if err := f.db.Find(&flights).Error; err != nil {
		return nil, err
	}
	return flights, nil
}

func (f *flightRepository) UpdateFlight(flight *models.FlightDetail) error {
	if err := f.db.Save(flight).Error; err != nil {
		return err
	}
	return nil
}

func (f *flightRepository) DeleteFlight(id uint) error {
	result := f.db.Delete(&models.FlightDetail{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ====== FlightService Methods ======

func (f *flightRepository) CreateFlightService(service *models.FlightService) error {
	if err := f.db.Create(&service).Error; err != nil {
		return err
	}
	return nil
}

func (f *flightRepository) GetAllFlightService() ([]*models.FlightService, error) {
	var service []*models.FlightService
	if err := f.db.Find(&service).Error; err != nil {
		return nil, err
	}
	return service, nil
}

func (f *flightRepository) UpdateFlightService(service *models.FlightService) error {
	if err := f.db.Save(service).Error; err != nil {
		return err
	}
	return nil
}

func (f *flightRepository) DeleteFlightService(id uint) error {
	result := f.db.Delete(&models.FlightService{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}


func (s *flightRepository) CheckIfSeatAvailable(seatID uint) (bool, error) {
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
func (s *flightRepository) GetSeatByID(seatID uint) (*models.SeatDetail, error) {
	var seat models.SeatDetail
	if err := s.db.First(&seat, seatID).Error; err != nil {
		return nil, err
	}
	return &seat, nil
}


// GetSeatsByFlightID implements SeatRepository.
func (s *flightRepository) GetSeatsByFlightID(flightID uint) ([]*models.SeatDetail, error) {
	var seats []*models.SeatDetail
	if err := s.db.Where("FlightID = ?", flightID).Find(&seats).Error; err != nil {
		return nil, err
	}
	return seats, nil
}


// ReleaseSeat implements SeatRepository.
func (s *flightRepository) ReleaseSeat(seatID uint) error {
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
func (s *flightRepository) ReserveSeat(seatID uint, reservationID uint) error {
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
func (s *flightRepository) UpdateSeat(seat *models.SeatDetail) error {
	if err := s.db.Save(seat).Error; err != nil {
		return err
	}
	return nil
}