package payment

import (
	"example.com/airline-reservation/pkg/container"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PaymentModule struct {
	E *echo.Echo
}

func (m *PaymentModule) Register(db *gorm.DB, c *container.AppContainer) {
	repo := NewPaymentRepository(db)
	service := NewPaymentService(repo)
	handler := NewPaymentHandler(service)

	c.UserService = service

	RegisterPaymentRoutes(m.E, handler)
}