package airport

import "example.com/airline-reservation/models"

type AirportService interface {
	GetAllAirports() ([]*models.Airport, error)
	CreateAirport(airport *models.Airport) error
	GetAirportByID(id uint) (*models.Airport, error)
	UpdateAirport(id uint, airport *models.Airport) (*models.Airport, error)
	DeleteAirport(id uint) error
}

type airportService struct {
	repo AirPortRepository
}

// GetAllAirport implements AirportService.


func NewAirportService(repo AirPortRepository) AirportService {
	return &airportService{repo: repo}
}

func (s *airportService) GetAllAirports() ([]*models.Airport, error) {
	return s.repo.GetAllAirports()
}

func (s *airportService) CreateAirport(airport *models.Airport) error {
	return s.repo.CreateAirPort(airport)
}

func (s *airportService) GetAirportByID(id uint) (*models.Airport, error) {
	return s.repo.GetaAirportById(id)
}

func (s *airportService) UpdateAirport(id uint, airport *models.Airport) (*models.Airport, error) {
	existingairport, err := s.repo.GetaAirportById(id)
	if err != nil {
		return nil, err
	}

	existingairport.Name = airport.Name
	existingairport.AirportCity = airport.AirportCity
	existingairport.AirportCountry = airport.AirportCountry
	existingairport.FlightsFrom = airport.FlightsFrom
	existingairport.FlightsTo = airport.FlightsTo

	err = s.repo.UpdateAirport(existingairport)
	if err != nil {
		return nil, err
	}
	return existingairport, nil
}

func (s *airportService) DeleteAirport(id uint) error {
	return s.repo.DeleteAirport(id)
}
