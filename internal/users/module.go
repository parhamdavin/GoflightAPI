package users

import (
	"example.com/airline-reservation/pkg/container"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserModule struct {
	E *echo.Echo
}

func (m *UserModule) Register(db *gorm.DB, c *container.AppContainer) {
	repo := NewUserRepository(db)
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	c.UserService = service

	RegisterUserRoutes(m.E, handler)
}
