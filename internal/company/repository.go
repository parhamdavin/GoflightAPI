package company
import "example.com/airline-reservation/models"

type CompanyRepository interface{
	CreateCompany(comapany *models.Company) error
	UpdateCompany(comapany *models.Company) error
	GetAllCompany()([]*models.Company,error)
	GetCompanyByID(id uint)(*models.Company,error)

}