package airport

import(
	"example.com/airline-reservation/models"
)

type AirPortRepository interface{
	CreateAirPort(airport *models.Airport) error
	GetaAirportById(id uint) (*models.Airport, error)
	UpdateAirport(airport *models.Airport)  error
	DeleteAirport(id uint) error
	GetAllAirports()([]*models.Airport, error)
	
}