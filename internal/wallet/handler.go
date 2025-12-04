package wallet

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"example.com/airline-reservation/models"
)

type WalletHandler struct {
	service WalletService
}

func NewWalletHandler(service WalletService) *WalletHandler {
	return &WalletHandler{service: service}
}


// POST /wallets
func (h *WalletHandler) CreateWallet(c echo.Context) error {
	var wallet models.Wallet
	if err := c.Bind(&wallet); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid Data"})
	}

	if err := h.service.CreateWallet(&wallet); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, wallet)
}

// GET /wallets/user/:id
func (h *WalletHandler) GetWalletByUserID(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "شناسه نامعتبر"})
	}

	wallet, err := h.service.GetWalletByUserID(uint(userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "کیف پول پیدا نشد"})
	}

	return c.JSON(http.StatusOK, wallet)
}

// PUT /wallets/balance
func (h *WalletHandler) UpdateWalletBalance(c echo.Context) error {
	var wallet models.Wallet
	if err := c.Bind(&wallet); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": " Invalid Data"})
	}

	if err := h.service.UpdateWalletBalance(&wallet); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, wallet)
}

// POST /wallets/transactions
func (h *WalletHandler) CreateWalletTransaction(c echo.Context) error {
	var tx models.WalletTransaction
	if err := c.Bind(&tx); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": " Invalid Data"})
	}

	if err := h.service.CreateWalletTransaction(&tx); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, tx)
}

// GET /wallets/:id/transactions
func (h *WalletHandler) GetTransactionsByWalletID(c echo.Context) error {
	walletID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "شناسه کیف پول نامعتبر"})
	}

	transactions, err := h.service.GetTransactionsByWalletID(uint(walletID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, transactions)
}
