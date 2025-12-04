package auth

// import (
// 	"example.com/airline-reservation/internal/users"
// 	"example.com/airline-reservation/pkg/container"
// 	"github.com/labstack/echo/v4"
// 	"gorm.io/gorm"
// )

// type AutModule struct {
// 	E *echo.Echo
// }

// func (m *AutModule) Register(db *gorm.DB, c *container.AppContainer) {
// 	// Repository و Service مربوط به User رو ایجاد می‌کنیم
// 	repo := users.NewUserRepository(db)
// 	service := users.NewUserService(repo)

// 	jwtService := NewJWTService("apikeys") 

// 	// AuthHandler رو با استفاده از Service و JWTService ایجاد می‌کنیم
// 	handler := NewAuthHandler(service, jwtService)

// 	// حالا به container اپلیکیشن سرویس‌های مربوطه رو می‌دیم
// 	c.UserService = service
// 	c.JWTService = jwtService

// 	// روت‌های مربوط به Auth رو ثبت می‌کنیم
// 	RegisterAuthRoutes(m.E, handler)
// }
