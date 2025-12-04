package reservations

import (
	"example.com/airline-reservation/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterReservationRoutes(e *echo.Echo, handler *ReservationHandler) {
	r := e.Group("/reservations")
	r.POST("/create", handler.CreateReservation)
	r.GET("/get/:id", handler.GetReservationByID)
	r.GET("/getreservationsbyuser/:userID", handler.GetReservationsByUser,middleware.JWTMiddleware)
}
