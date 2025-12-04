package wallet

import (
	"example.com/airline-reservation/pkg/container"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type WalletModule struct {
	E *echo.Echo
}

func (m *WalletModule) Register(db *gorm.DB, c *container.AppContainer) {
	repo := NewWalletRepository(db)
	service := NewWalletService(repo)
	handler := NewWalletHandler(service)

	c.UserService = service

	RegisterWalletRoutes(m.E, handler)
}