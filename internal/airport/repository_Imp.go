package airport

import (
	"example.com/airline-reservation/models"
	"gorm.io/gorm"
)

type airportRepository struct {
	db *gorm.DB
}





func NewAirportRepository(db *gorm.DB) AirPortRepository {
	return &airportRepository{db}
}

func (a *airportRepository) GetAllAirports() ([]*models.Airport, error) {
	var airports []*models.Airport
	if err := a.db.Find(&airports).Error; err != nil {
		return nil, err
	}
	return airports, nil
}

func (a *airportRepository) CreateAirPort(airport *models.Airport) error {
    if err := a.db.Create(&airport).Error; err != nil {
        return err
    }
    return nil
}

func (a *airportRepository) DeleteAirport(id uint) error {
	result := a.db.Delete(&models.Airport{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}


func (a *airportRepository) GetaAirportById(id uint) (*models.Airport, error) {
	airport := &models.Airport{}

	if err := a.db.First(airport, id).Error; err != nil {
		return nil, err 
	}

	return airport, nil 
}


func (a *airportRepository) UpdateAirport(airport *models.Airport)  error {
	if err := a.db.Save(airport).Error; err != nil {
		return  err
	}
	return  nil
}

