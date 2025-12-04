package users

import (
	"github.com/labstack/echo/v4"
	"example.com/airline-reservation/internal/middleware"
)

func RegisterUserRoutes(e *echo.Echo, handler *UserHandler) {
	users := e.Group("/users")
	users.POST("/register", handler.Register) // بدون نیاز به توکن
	users.POST("/login", handler.Login)       // بدون نیاز به توکن

	// اینا نیاز به توکن دارن
	users.GET("", handler.GetAllUsers, middleware.JWTMiddleware)
	users.POST("", handler.CreateUser, middleware.JWTMiddleware)
	users.GET("/:id", handler.GetUserByID, middleware.JWTMiddleware)
	users.PUT("/:id", handler.UpdateUser, middleware.JWTMiddleware)
	users.DELETE("/:id", handler.DeleteUser, middleware.JWTMiddleware)
}

