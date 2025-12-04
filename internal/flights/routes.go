package flights

import (
	"github.com/labstack/echo/v4"
	"example.com/airline-reservation/internal/middleware"
)

func RegisterFlightRoutes(e *echo.Echo, handler *FlightHandler) {
	flights := e.Group("/flights")

	// FlightDetail routes
	flights.GET("/getall", handler.GetAllFlights)                                    
	flights.GET("/get/:id", handler.GetFlightByID)                                
	flights.POST("/create", handler.CreateFlight, middleware.JWTMiddleware)           // نیاز به توکن
	flights.PUT("/update/:id", handler.UpdateFlight, middleware.JWTMiddleware)        // نیاز به توکن
	flights.DELETE("/delete/:id", handler.DeleteFlight, middleware.JWTMiddleware)     // نیاز به توکن

	// FlightService routes
	flightServices := e.Group("/flight-services")
	flightServices.GET("/getall", handler.GetAllFlightServices)
	flightServices.GET("/get/:id", handler.GetFlightServiceByID)
	flightServices.POST("/create", handler.CreateFlightService, middleware.JWTMiddleware)
	flightServices.PUT("/update/:id", handler.UpdateFlightService, middleware.JWTMiddleware)
	flightServices.DELETE("/delete/:id", handler.DeleteFlightService, middleware.JWTMiddleware)

	// FlightCost routes
	flightCosts := e.Group("/flight-costs")
	flightCosts.GET("/getall", handler.GetAllFlightCosts)
	flightCosts.GET("/get/:id", handler.GetFlightCostByID)
	flightCosts.POST("/create", handler.CreateFlightCost, middleware.JWTMiddleware)
	flightCosts.PUT("/update/:id", handler.UpdateFlightCost, middleware.JWTMiddleware)
	flightCosts.DELETE("/delete/:id", handler.DeleteFlightCost, middleware.JWTMiddleware)

	flightcalendar := e.Group("/flight-calendar")
	flightcalendar.GET("/getall", handler.GetAllCalendars)
	flightcalendar.GET("/get/:id", handler.GetCalendarByID)
	flightcalendar.POST("/create", handler.CreateCalendar, middleware.JWTMiddleware)
	flightcalendar.PUT("/update/:id", handler.UpdateCalendar, middleware.JWTMiddleware)
	flightcalendar.DELETE("/delete/:id", handler.DeleteCalendar, middleware.JWTMiddleware)


	flightseat := e.Group("/flight-seat")
	
	flightseat.GET("/get/:id", handler.GetSeatByID)
	flightseat.GET("/get-by-flight/:id", handler.GetSeatsByFlightID)
	flightseat.POST("/release-seat/:id", handler.ReleaseSeat,)
	flightseat.PUT("/update/:id", handler.UpdateSeat, middleware.JWTMiddleware)
	flightseat.POST("/reserve-seat/:id", handler.ReserveSeat)
}
