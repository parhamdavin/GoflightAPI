package models

import (
	"time"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	UserId uint `gorm:"uniqueIndex"`
	Balance   float64
	Transactions []WalletTransaction `gorm:"foreignKey:WalletID"`
}

type WalletTransaction struct {
	ID          uint
	WalletID    uint           
	Amount      float64
	Type        string         
	Status      string         
	Description string
	CreatedAt   time.Time
}
