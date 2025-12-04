package wallet

import "example.com/airline-reservation/models"

type WalletService interface {
	CreateWallet(wallet *models.Wallet) error
	GetWalletByUserID(UserId uint) (*models.Wallet, error)
	UpdateWalletBalance(wallet *models.Wallet) error
	CreateWalletTransaction(Transaction *models.WalletTransaction) error
	GetTransactionsByWalletID(walletID uint) ([]*models.WalletTransaction, error)
}

type walletservice struct {
	repo WalletRepository
}

// CreateWallet implements WalletService.
func (w *walletservice) CreateWallet(wallet *models.Wallet) error {
	return w.repo.CreateWallet(wallet)
}

// CreateWalletTransaction implements WalletService.
func (w *walletservice) CreateWalletTransaction(Transaction *models.WalletTransaction) error {
	return w.repo.CreateWalletTransaction(Transaction)
}

// GetTransactionsByWalletID implements WalletService.
func (w *walletservice) GetTransactionsByWalletID(walletID uint) ([]*models.WalletTransaction, error) {
	return w.repo.GetTransactionsByWalletID(walletID)
}

// GetWalletByUserID implements WalletService.
func (w *walletservice) GetWalletByUserID(UserId uint) (*models.Wallet, error) {
	return w.repo.GetWalletByUserID(UserId)
}

// UpdateWalletBalance implements WalletService.
func (w *walletservice) UpdateWalletBalance(wallet *models.Wallet) error {
	return w.repo.UpdateWalletBalance(wallet)
}

func NewWalletService(repo WalletRepository) WalletService {
	return &walletservice{repo: repo}
}
