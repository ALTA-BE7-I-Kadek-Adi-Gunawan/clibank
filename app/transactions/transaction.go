package transactions

import (
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/users"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Type       string         `gorm:"type:enum('Debit','Credit');not null"`
	ReceiverID uint           `json:"-"`
	SenderID   uint           `json:"-"`
	Ammount    float64        `gorm:"type:decimal(10,2);not null"`
	Receiver   *users.Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Sender     *users.Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
