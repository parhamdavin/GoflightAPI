package company

import (
	"example.com/airline-reservation/models"
	"gorm.io/gorm"
)

type companyrepository struct {
	db *gorm.DB
}

// GetCompanyByID implements CompanyRepository.
func (c *companyrepository) GetCompanyByID(id uint) (*models.Company, error) {
	company:=&models.Company{}
	if err := c.db.First(company, id).Error; err != nil {
		return nil, err
	}

	return company,nil
}

func (c *companyrepository) CreateCompany(comapany *models.Company) error {
	if err := c.db.Create(comapany).Error; err != nil {
		return err
	}
	return nil
}

// GetAllCompany implements CompanyRepository.
func (c *companyrepository) GetAllCompany() ([]*models.Company, error) {
	var company []*models.Company
	if err := c.db.Find(&company).Error; err != nil {
		return nil, err
	}
	return company, nil
}

// UpdateCompany implements CompanyRepository.
func (c *companyrepository) UpdateCompany(comapany *models.Company) error {
	if err := c.db.Save(comapany).Error; err != nil {
		return err
	}
	return nil
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyrepository{db}
}
