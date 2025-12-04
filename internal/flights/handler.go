package flights

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"example.com/airline-reservation/models"
)

type FlightHandler struct {
	Service FlightService
}

func NewFlightHandler(service FlightService) *FlightHandler {
	return &FlightHandler{Service: service}
}

// ---------- FlightDetail Handlers ----------

func (h *FlightHandler) CreateFlight(c echo.Context) error {
	var flight models.FlightDetail
	if err := c.Bind(&flight); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := h.Service.CreateFlight(&flight); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, flight)
}

func (h *FlightHandler) GetAllFlights(c echo.Context) error {
	flights, err := h.Service.GetAllFlight()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, flights)
}

func (h *FlightHandler) GetFlightByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	flight, err := h.Service.GetFlightByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, flight)
}

func (h *FlightHandler) UpdateFlight(c echo.Context) error {
	var flight models.FlightDetail
	if err := c.Bind(&flight); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := h.Service.UpdateFlight(&flight); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, flight)
}

func (h *FlightHandler) DeleteFlight(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteFlight(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// ---------- FlightService Handlers ----------

func (h *FlightHandler) CreateFlightService(c echo.Context) error {
	var service models.FlightService
	if err := c.Bind(&service); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := h.Service.CreateFlightService(&service); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, service)
}

func (h *FlightHandler) GetAllFlightServices(c echo.Context) error {
	services, err := h.Service.GetAllFlightService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, services)
}

func (h *FlightHandler) GetFlightServiceByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	service, err := h.Service.GetFlightServiceByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (h *FlightHandler) UpdateFlightService(c echo.Context) error {
	var service models.FlightService
	if err := c.Bind(&service); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := h.Service.UpdateFlightService(&service); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (h *FlightHandler) DeleteFlightService(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteFlightService(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// ---------- FlightCost Handlers ----------

func (h *FlightHandler) CreateFlightCost(c echo.Context) error {
	var cost models.FlightCost
	if err := c.Bind(&cost); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := h.Service.CreateFlightCost(&cost); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, cost)
}

func (h *FlightHandler) GetAllFlightCosts(c echo.Context) error {
	costs, err := h.Service.GetAllFlightCost()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, costs)
}

func (h *FlightHandler) GetFlightCostByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	cost, err := h.Service.GetFlightCostByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, cost)
}

func (h *FlightHandler) UpdateFlightCost(c echo.Context) error {
	var cost models.FlightCost
	if err := c.Bind(&cost); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	if err := h.Service.UpdateFlightCost(&cost); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, cost)
}

func (h *FlightHandler) DeleteFlightCost(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteFlightCost(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}


func (h *FlightHandler) CreateCalendar(c echo.Context) error {
	var calendar models.Calendar
	if err := c.Bind(&calendar); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "فرمت ورودی نادرست است"})
	}

	if err := h.Service.CreateCalendar(&calendar); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, calendar)
}

func (h *FlightHandler) DeleteCalendar(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "شناسه نامعتبر است"})
	}

	if err := h.Service.DeleteCalendar(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *FlightHandler) GetAllCalendars(c echo.Context) error {
	calendars, err := h.Service.GetAllCalendar()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, calendars)
}

func (h *FlightHandler) GetCalendarByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "شناسه نامعتبر است"})
	}

	calendar, err := h.Service.GetCalendarByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "تاریخ پیدا نشد"})
	}

	return c.JSON(http.StatusOK, calendar)
}

func (h *FlightHandler) UpdateCalendar(c echo.Context) error {
	var calendar models.Calendar
	if err := c.Bind(&calendar); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "فرمت ورودی نادرست است"})
	}

	if err := h.Service.UpdateCalendar(&calendar); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, calendar)
}


// GET /seats/:id
func (h *FlightHandler) GetSeatByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "شناسه صندلی نامعتبر است"})
	}

	seat, err := h.Service.GetSeatByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "صندلی پیدا نشد"})
	}

	return c.JSON(http.StatusOK, seat)
}

// GET /flights/:flight_id/seats
func (h *FlightHandler) GetSeatsByFlightID(c echo.Context) error {
	flightID, err := strconv.Atoi(c.Param("flight_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "شناسه پرواز نامعتبر است"})
	}

	seats, err := h.Service.GetSeatsByFlightID(uint(flightID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, seats)
}

// GET /seats/:id/available
func (h *FlightHandler) CheckIfSeatAvailable(c echo.Context) error {
	seatID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "شناسه صندلی نامعتبر است"})
	}

	available, err := h.Service.CheckIfSeatAvailable(uint(seatID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"available": available})
}

// PUT /seats/:id/release
func (h *FlightHandler) ReleaseSeat(c echo.Context) error {
	seatID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "شناسه صندلی نامعتبر است"})
	}

	err = h.Service.ReleaseSeat(uint(seatID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.NoContent(http.StatusOK)
}

// PUT /seats/:id/reserve/:reservation_id
func (h *FlightHandler) ReserveSeat(c echo.Context) error {
	seatID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "شناسه صندلی نامعتبر است"})
	}

	reservationID, err := strconv.Atoi(c.Param("reservation_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "شناسه رزرو نامعتبر است"})
	}

	err = h.Service.ReserveSeat(uint(seatID), uint(reservationID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.NoContent(http.StatusOK)
}

// PUT /seats/:id
func (h *FlightHandler) UpdateSeat(c echo.Context) error {
	var seat models.SeatDetail
	if err := c.Bind(&seat); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "داده‌های ورودی نامعتبر است"})
	}

	err := h.Service.UpdateSeat(&seat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.NoContent(http.StatusOK)
}
