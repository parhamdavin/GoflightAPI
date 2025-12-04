package users

import (
	"example.com/airline-reservation/models"
)

type UserService interface {
	GetAllUsers() ([]*models.User, error)
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(id uint, user *models.User) (*models.User, error)
	DeleteUser(id uint) error
	GetByEmail(email string) (*models.User, error)
}

type userService struct {
	repo UserRepository
}

// GetByEmail implements UserService.
func (s *userService) GetByEmail(email string) (*models.User, error) {
	return s.repo.GetByEmail(email)
}


func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]*models.User, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) CreateUser(user *models.User) error {
	return s.repo.CraeteUser(user)
}

func (s *userService) GetUserByID(id uint) (*models.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(id uint, user *models.User) (*models.User, error) {
	existingUser, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	
	existingUser.Name = user.Name
	existingUser.SSID = user.SSID
	existingUser.PhoneNumber = user.PhoneNumber
	existingUser.PassportId = user.PassportId
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	existingUser.Agency = user.Agency
	existingUser.Active = user.Active
	existingUser.Role = user.Role

	err = s.repo.UpdateUser(existingUser)
	if err != nil {
		return nil, err
	}
	return existingUser, nil
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
