package auth

// import (
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// )

// func RegisterAuthRoutes(e *echo.Echo, h *AuthHandler) {
// 	auth := e.Group("/auth")
// 	auth.POST("/login", h.Login)
// 	auth.POST("/register", h.Register)

// 	// مثلا مسیرهای محافظت شده:
// 	protected := e.Group("/protected")
// 	protected.Use(JWTMiddleware(h.JWTService)) // 👈 میدلور اضافه می‌کنیم
// 	protected.GET("/profile", func(c echo.Context) error {
// 		return c.JSON(http.StatusOK, echo.Map{"message": "you are authorized!"})
// 	})
// }

