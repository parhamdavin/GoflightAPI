package reservations



import (
	"example.com/airline-reservation/pkg/container"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ReservationModule struct {
	E *echo.Echo
}

func (m *ReservationModule) Register(db *gorm.DB, c *container.AppContainer) {
	repo := NewReservationRepository(db)
	service := NewReservationService(repo)
	handler := NewReservationHandler(service)

	c.UserService = service

	RegisterReservationRoutes(m.E, handler)
}