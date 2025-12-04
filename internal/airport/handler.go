package airport

import (
	"net/http"
	"strconv"

	"example.com/airline-reservation/models"
	"github.com/labstack/echo/v4"
)

type AirportHandler struct {
	Service AirportService
}

func NewAirportHandler(service AirportService) *AirportHandler {
	return &AirportHandler{Service: service}
}

// گرفتن همه فرودگاه‌ها
func (h *AirportHandler) GetAllAirports(c echo.Context) error {
	airports, err := h.Service.GetAllAirports()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, airports)
}

// ساخت فرودگاه جدید
func (h *AirportHandler) CreateAirport(c echo.Context) error {
	var airportInput models.Airport
	if err := c.Bind(&airportInput); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	err := h.Service.CreateAirport(&airportInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, airportInput)
}

// گرفتن فرودگاه با آیدی
func (h *AirportHandler) GetAirportByID(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid airport ID"})
	}

	airport, err := h.Service.GetAirportByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "airport not found"})
	}
	return c.JSON(http.StatusOK, airport)
}

// آپدیت فرودگاه
func (h *AirportHandler) UpdateAirport(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid airport ID"})
	}

	var airportInput models.Airport
	if err := c.Bind(&airportInput); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	updatedAirport, err := h.Service.UpdateAirport(uint(id), &airportInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedAirport)
}

// حذف فرودگاه
func (h *AirportHandler) DeleteAirport(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid airport ID"})
	}

	err = h.Service.DeleteAirport(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
