package wallet

import "example.com/airline-reservation/models"

type WalletRepository interface {
	CreateWallet(wallet *models.Wallet)error
	GetWalletByUserID(UserId uint)(*models.Wallet , error)
	UpdateWalletBalance(wallet *models.Wallet)error
	CreateWalletTransaction(Transaction *models.WalletTransaction)error
	GetTransactionsByWalletID(walletID uint) ([]*models.WalletTransaction, error)
}