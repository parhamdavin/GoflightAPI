 package auth

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/labstack/echo/v4"
// )

// func JWTMiddleware(jwtService JWTService) echo.MiddlewareFunc {
// 	return func(next echo.HandlerFunc) echo.HandlerFunc {
// 		return func(c echo.Context) error {
// 			authHeader := c.Request().Header.Get("Authorization")
// 			if authHeader == "" {
// 				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing or invalid token"})
// 			}

// 			tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

// 			token, err := jwtService.ValidateToken(tokenString)
// 			if err != nil {
// 				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid or expired token"})
// 			}

// 			if !token.Valid {
// 				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
// 			}

// 			return next(c)
// 		}
// 	}
// }
