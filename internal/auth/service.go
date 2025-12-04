package auth

// import (
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

// // JWTService defines the methods for token generation and validation.
// type JWTService interface {
// 	GenerateToken(userID uint, email, role string) (string, error)
// 	ValidateToken(token string) (*jwt.Token, error)
// 	GetSecret() string
// }

// type jwtService struct {
// 	secretKey string
// }

// // NewJWTService creates a new instance of JWTService.
// func NewJWTService(secret string) JWTService {
// 	return &jwtService{secretKey: secret}
// }

// // GetSecret returns the secret key.
// func (j *jwtService) GetSecret() string {
// 	return j.secretKey
// }

// // GenerateToken creates a JWT token for the user.
// func (j *jwtService) GenerateToken(userID uint, email, role string) (string, error) {
// 	claims := jwt.MapClaims{
// 		"user_id": userID,
// 		"email":   email,
// 		"role":    role,
// 		"exp":     time.Now().Add(time.Hour * 24).Unix(),
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(j.secretKey))
// }

// // ValidateToken parses and validates a JWT token.
// func (j *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
// 	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(j.secretKey), nil
// 	})
// }
