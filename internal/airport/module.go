package airport


import (
	"example.com/airline-reservation/pkg/container"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AirportModule struct {
	E *echo.Echo
}

func (m *AirportModule) Register(db *gorm.DB, c *container.AppContainer) {
	repo := NewAirportRepository(db)
	service := NewAirportService(repo)
	handler := NewAirportHandler(service)

	c.AirportService = service

	RegisterAirportRoutes(m.E, handler)
}
