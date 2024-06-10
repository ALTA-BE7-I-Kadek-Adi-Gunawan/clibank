package topups

import (
	"github.com/1layar/clibank/app/wallets"
	"gorm.io/gorm"
)

type TopupWallet struct {
	gorm.Model
	Amount        float64        `gorm:"type:decimal(10,2);not null"`
	SerialNo      string         `gorm:"type:varchar(100);not null;uniqueIndex"`
	Status        string         `gorm:"type:enum('pending','success','failed');default:'success'"`
	WalletID      uint           `json:"-"`
	TopupOptionID uint           `json:"-"`
	Wallet        wallets.Wallet `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TopupOption   TopupOption    `gorm:"foreignKey:TopupOptionID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
