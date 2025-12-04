package users

import (
	"fmt"

	"example.com/airline-reservation/models"
	"gorm.io/gorm"
)

type userrepository struct {
	db *gorm.DB
}

// CraeteUser implements UserRepository.
func (u *userrepository) CraeteUser(user *models.User) error {
	fmt.Println("🔥 Inside CraeteUser - password being saved:", user.Password)

	if err := u.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}


// DeleteUser implements UserRepository.
func (u *userrepository) DeleteUser(id uint) error {
	result := u.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllUsers implements UserRepository.
func (u *userrepository) GetAllUsers() ([]*models.User, error) {
	var user []*models.User
	if err := u.db.Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByID implements UserRepository.
func (u *userrepository) GetUserByID(id uint) (*models.User, error) {
	user := &models.User{}

	if err := u.db.First(user, id).Error; err != nil {
		return nil, err
	}

	return user,nil
}

// UpdateUser implements UserRepository.
func (u *userrepository) UpdateUser(user *models.User) error {
	if err := u.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}
// GetByEmail implements UserRepository.
func (u *userrepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := u.db.Preload("Wallet").Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userrepository{db}
}
