 package auth

// import (
	
// 	"net/http"

// 	"example.com/airline-reservation/internal/users"
// 	"example.com/airline-reservation/models"
// 	"github.com/labstack/echo/v4"
// 	"golang.org/x/crypto/bcrypt"
// )

// type AuthHandler struct {
// 	UserService users.UserService // از لایه users
// 	JWTService  JWTService
// }

// func NewAuthHandler(us users.UserService, jwt JWTService) *AuthHandler {
// 	return &AuthHandler{
// 		UserService: us,
// 		JWTService:  jwt,
// 	}
// }

// type LoginRequest struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// func (h *AuthHandler) Login(c echo.Context) error {
// 	var req LoginRequest
// 	if err := c.Bind(&req); err != nil {
// 		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input", "details": err.Error()})
// 	}

// 	user, err := h.UserService.GetByEmail(req.Email)
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, echo.Map{
// 			"error":   "user not found",
// 			"details": err.Error(),
// 		})
// 	}
	
// 	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
// 		return c.JSON(http.StatusUnauthorized, echo.Map{
// 			"error":   "password mismatch",
// 			"details": err.Error(),
// 		})
// 	}

// 	token, err := h.JWTService.GenerateToken(user.ID, user.Email, user.Role)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, echo.Map{
// 			"error":   "failed to generate token",
// 			"details": err.Error(),
// 		})
// 	}

// 	return c.JSON(http.StatusOK, echo.Map{
// 		"token": token,
// 		"user":  user,
// 	})
// }

// type RegisterRequest struct {
// 	Name         string `json:"name"`
// 	Email        string `json:"email"`
// 	Password     string `json:"password"`
// 	PhoneNumber  string `json:"phone_number"`
// 	SSID         string `json:"ssid"`
// 	PassportId   string `json:"passport_id"`
// 	Role         string `json:"role"`
// 	Agency       bool   `json:"agency"`
// 	Active       bool   `json:"active"`
// }

// func (h *AuthHandler) Register(c echo.Context) error {
// 	var req RegisterRequest
// 	if err := c.Bind(&req); err != nil {
// 		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
// 	}

// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to hash password"})
// 	}

// 	user := &models.User{
// 		Name:        req.Name,
// 		Email:       req.Email,
// 		Password:    string(hashedPassword),
// 		PhoneNumber: req.PhoneNumber,
// 		SSID:        req.SSID,
// 		PassportId:  req.PassportId,
// 		Role:        req.Role,
// 		Agency:      req.Agency,
// 		Active:      req.Active,
// 	}

// 	err = h.UserService.CreateUser(user)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
// 	}

// 	return c.JSON(http.StatusCreated, echo.Map{"message": "user registered successfully", "user": user})
// }
