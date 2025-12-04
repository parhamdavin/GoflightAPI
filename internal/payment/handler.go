package payment

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"example.com/airline-reservation/models"
)

type PaymentHandler struct {
	service PaymentService
}

func NewPaymentHandler(service PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

// POST /payments
func (h *PaymentHandler) CreatePayment(c echo.Context) error {
	var payment models.PaymentStatus
	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "داده‌های نامعتبر"})
	}

	if err := h.service.CraetePayment(&payment); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, payment)
}

// GET /payments/:id
func (h *PaymentHandler) GetPaymentByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "شناسه نامعتبر"})
	}

	payment, err := h.service.GetPaymentByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "پرداخت پیدا نشد"})
	}

	return c.JSON(http.StatusOK, payment)
}

// GET /reservations/:id/payments
func (h *PaymentHandler) GetPaymentsByReservation(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "شناسه رزرو نامعتبر است"})
	}

	payments, err := h.service.GetPaymentByReservation(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, payments)
}
