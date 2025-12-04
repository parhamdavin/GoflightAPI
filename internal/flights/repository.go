package flights

import "example.com/airline-reservation/models"

type FlightRespository interface {
	CreateFlight(flight *models.FlightDetail)error
	UpdateFlight(flight *models.FlightDetail) error
	DeleteFlight(id uint) error
	GetAllFlight()([]*models.FlightDetail,error)
	GetFlightByID(id uint)(*models.FlightDetail,error)
	CreateFlightService(service *models.FlightService)error
	UpdateFlightService(service *models.FlightService) error
	DeleteFlightService(id uint) error
	GetAllFlightService()([]*models.FlightService,error)
	GetFlightServiceByID(id uint)(*models.FlightService,error)
	CreateFlightCost(cost *models.FlightCost)error
	UpdateFlightCost(cost *models.FlightCost) error
	DeleteFlightCost(id uint) error
	GetAllFlightCost()([]*models.FlightCost,error)
	GetFlightCostByID(id uint)(*models.FlightCost,error)
	CreateCalendar(calendar *models.Calendar)error
	UpdateCalendar(calendar *models.Calendar) error
	DeleteCalendar(id uint) error
	GetAllCalendar()([]*models.Calendar,error)
	GetCalendarByID(id uint)(*models.Calendar,error)

	GetSeatsByFlightID(flightID uint) ([]*models.SeatDetail, error)
	ReserveSeat(seatID uint, reservationID uint) error
	ReleaseSeat(seatID uint) error
	GetSeatByID(seatID uint) (*models.SeatDetail, error)
	UpdateSeat(seat *models.SeatDetail) error
	CheckIfSeatAvailable(seatID uint) (bool, error)
	
}