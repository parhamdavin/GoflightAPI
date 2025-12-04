package airport

import "github.com/labstack/echo/v4"

func RegisterAirportRoutes(e *echo.Echo, handler *AirportHandler) {
	airports := e.Group("/airports")

	airports.GET("", handler.GetAllAirports)
	airports.GET("/:id", handler.GetAirportByID)
	airports.POST("/create", handler.CreateAirport)
	airports.PUT("/:id", handler.UpdateAirport)
	airports.DELETE("/:id", handler.DeleteAirport)
}
