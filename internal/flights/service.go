package flights

import (
	"errors"
	"example.com/airline-reservation/models"
)

type FlightService interface {
	CreateFlight(flight *models.FlightDetail) error
	UpdateFlight(flight *models.FlightDetail) error
	DeleteFlight(id uint) error
	GetAllFlight() ([]*models.FlightDetail, error)
	GetFlightByID(id uint) (*models.FlightDetail, error)

	CreateFlightService(service *models.FlightService) error
	UpdateFlightService(service *models.FlightService) error
	DeleteFlightService(id uint) error
	GetAllFlightService() ([]*models.FlightService, error)
	GetFlightServiceByID(id uint) (*models.FlightService, error)

	CreateFlightCost(cost *models.FlightCost) error
	UpdateFlightCost(cost *models.FlightCost) error
	DeleteFlightCost(id uint) error
	GetAllFlightCost() ([]*models.FlightCost, error)
	GetFlightCostByID(id uint) (*models.FlightCost, error)

	CreateCalendar(calendar *models.Calendar) error
	UpdateCalendar(calendar *models.Calendar) error
	DeleteCalendar(id uint) error
	GetAllCalendar() ([]*models.Calendar, error)
	GetCalendarByID(id uint) (*models.Calendar, error)

	GetSeatsByFlightID(flightID uint) ([]*models.SeatDetail, error)
	ReserveSeat(seatID uint, reservationID uint) error
	ReleaseSeat(seatID uint) error
	GetSeatByID(seatID uint) (*models.SeatDetail, error)
	UpdateSeat(seat *models.SeatDetail) error
	CheckIfSeatAvailable(seatID uint) (bool, error)
}

type flightservice struct {
	repo FlightRespository
}

// CheckIfSeatAvailable implements FlightService.
func (f *flightservice) CheckIfSeatAvailable(seatID uint) (bool, error) {
	return f.repo.CheckIfSeatAvailable(seatID)
}

// GetSeatByID implements FlightService.
func (f *flightservice) GetSeatByID(seatID uint) (*models.SeatDetail, error) {
	return f.repo.GetSeatByID(seatID)
}

// GetSeatsByFlightID implements FlightService.
func (f *flightservice) GetSeatsByFlightID(flightID uint) ([]*models.SeatDetail, error) {
	return f.repo.GetSeatsByFlightID(flightID)
}

// ReleaseSeat implements FlightService.
func (f *flightservice) ReleaseSeat(seatID uint) error {
	return f.repo.ReleaseSeat(seatID)
}

// ReserveSeat implements FlightService.
func (f *flightservice) ReserveSeat(seatID uint, reservationID uint) error {
	return f.repo.ReserveSeat(seatID,reservationID)
}

// UpdateSeat implements FlightService.
func (f *flightservice) UpdateSeat(seat *models.SeatDetail) error {
	return f.repo.UpdateSeat(seat)
}

// CreateCalendar implements FlightService.
func (f *flightservice) CreateCalendar(calendar *models.Calendar) error {
	if calendar.Date.IsZero() {
		return errors.New("اطلاعات تاریخ ناقص است")
	}
	return f.repo.CreateCalendar(calendar)

}

// DeleteCalendar implements FlightService.
func (f *flightservice) DeleteCalendar(id uint) error {
	return f.repo.DeleteCalendar(id)
}

// GetAllCalendar implements FlightService.
func (f *flightservice) GetAllCalendar() ([]*models.Calendar, error) {
	return f.repo.GetAllCalendar()
}

// GetCalendarByID implements FlightService.
func (f *flightservice) GetCalendarByID(id uint) (*models.Calendar, error) {
	return f.repo.GetCalendarByID(id)
}

// UpdateCalendar implements FlightService.
func (f *flightservice) UpdateCalendar(calendar *models.Calendar) error {
	return f.repo.UpdateCalendar(calendar)
}

func NewFlightService(repo FlightRespository) FlightService {
	return &flightservice{repo: repo}
}

// ---------- FlightDetail ----------

func (f *flightservice) CreateFlight(flight *models.FlightDetail) error {
	if flight.FlightNumber == "" {
		return errors.New("اطلاعات پرواز ناقص است")
	}
	
	return f.repo.CreateFlight(flight)
}

func (f *flightservice) UpdateFlight(flight *models.FlightDetail) error {
	return f.repo.UpdateFlight(flight)
}

func (f *flightservice) DeleteFlight(id uint) error {
	return f.repo.DeleteFlight(id)
}

func (f *flightservice) GetAllFlight() ([]*models.FlightDetail, error) {
	return f.repo.GetAllFlight()
}

func (f *flightservice) GetFlightByID(id uint) (*models.FlightDetail, error) {
	return f.repo.GetFlightByID(id)
}

// ---------- FlightService ----------

func (f *flightservice) CreateFlightService(service *models.FlightService) error {
	return f.repo.CreateFlightService(service)
}

func (f *flightservice) UpdateFlightService(service *models.FlightService) error {
	return f.repo.UpdateFlightService(service)
}

func (f *flightservice) DeleteFlightService(id uint) error {
	return f.repo.DeleteFlightService(id)
}

func (f *flightservice) GetAllFlightService() ([]*models.FlightService, error) {
	return f.repo.GetAllFlightService()
}

func (f *flightservice) GetFlightServiceByID(id uint) (*models.FlightService, error) {
	return f.repo.GetFlightServiceByID(id)
}

// ---------- FlightCost ----------

func (f *flightservice) CreateFlightCost(cost *models.FlightCost) error {
	if cost.Cost <= 0 {
		return errors.New("هزینه پرواز باید بیشتر از صفر باشد")
	}
	calnder, err := f.repo.GetCalendarByID(cost.CalendarID)
	if err != nil {
		return errors.New("   زمان پرواز درست نیست   ")
	}
	cost.Calendar = *calnder

	return f.repo.CreateFlightCost(cost)
}

func (f *flightservice) UpdateFlightCost(cost *models.FlightCost) error {
	return f.repo.UpdateFlightCost(cost)
}

func (f *flightservice) DeleteFlightCost(id uint) error {
	return f.repo.DeleteFlightCost(id)
}

func (f *flightservice) GetAllFlightCost() ([]*models.FlightCost, error) {
	return f.repo.GetAllFlightCost()
}

func (f *flightservice) GetFlightCostByID(id uint) (*models.FlightCost, error) {
	return f.repo.GetFlightCostByID(id)
}
