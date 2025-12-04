package users

import (
	"example.com/airline-reservation/models"
	
)


type UserRepository interface{
	CraeteUser(user *models.User)error 
	DeleteUser(id uint)error
	UpdateUser(user *models.User)error
	GetAllUsers()([]*models.User,error)
	GetUserByID(id uint)(*models.User,error)
	GetByEmail(email string) (*models.User, error)

}