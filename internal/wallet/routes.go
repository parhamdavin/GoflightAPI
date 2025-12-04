package wallet
import (
	"example.com/airline-reservation/internal/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterWalletRoutes(e *echo.Echo, walletHandler *WalletHandler) {
	wallets := e.Group("/wallets")
	wallets.POST("", walletHandler.CreateWallet,middleware.JWTMiddleware)
	wallets.GET("/user/:id", walletHandler.GetWalletByUserID,middleware.JWTMiddleware)
	wallets.PUT("/balance", walletHandler.UpdateWalletBalance,middleware.JWTMiddleware)
	wallets.POST("/transactions", walletHandler.CreateWalletTransaction,middleware.JWTMiddleware)
	wallets.GET("/:id/transactions", walletHandler.GetTransactionsByWalletID,middleware.JWTMiddleware)

}