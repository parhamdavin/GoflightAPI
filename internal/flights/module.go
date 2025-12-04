package flights


import (
	"example.com/airline-reservation/pkg/container"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type FlightModule struct {
	E *echo.Echo
}

func (m *FlightModule) Register(db *gorm.DB, c *container.AppContainer) {
	repo := NewFlightRepository(db)
	service := NewFlightService(repo)
	handler := NewFlightHandler(service)

	c.UserService = service

	RegisterFlightRoutes(m.E, handler)
}