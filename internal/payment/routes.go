package payment

import (
	"github.com/labstack/echo/v4"
	
)

func RegisterPaymentRoutes(e *echo.Echo, handler *PaymentHandler) {
	payments := e.Group("/payments")
	payments.POST("", handler.CreatePayment)
	payments.GET("/:id", handler.GetPaymentByID)
	e.GET("/reservations/:id/payments", handler.GetPaymentsByReservation)

}