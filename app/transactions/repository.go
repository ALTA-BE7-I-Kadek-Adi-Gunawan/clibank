package transactions

import (
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/users"
	"gorm.io/gorm"
)

type ITransactionRepository interface {
	Transfer(sender *users.Account, receiver *users.Account, ammount float64) error
	TransferList() ([]Transaction, error)
}

type TransactionRepository struct {
	db *gorm.DB
}

func (u *TransactionRepository) Init(db *gorm.DB) {
	u.db = db
}

func (u *TransactionRepository) Transfer(sender *users.Account, receiver *users.Account, ammount float64) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		receiverTransaction := &Transaction{
			ReceiverID: receiver.ID,
			SenderID:   sender.ID,
			Type:       "Debit",
			Ammount:    ammount,
		}
		senderTransaction := &Transaction{
			ReceiverID: receiver.ID,
			SenderID:   sender.ID,
			Type:       "Credit",
			Ammount:    ammount,
		}
		receiverWallet := receiver.Wallet
		receiverWallet.Balance = receiverWallet.Balance + ammount
		senderWallet := sender.Wallet
		senderWallet.Balance = senderWallet.Balance - ammount
		err := tx.Save(receiverTransaction).Error
		if err != nil {
			return err
		}
		err = tx.Save(senderTransaction).Error
		if err != nil {
			return err

		}
		err = tx.Save(receiverWallet).Error
		if err != nil {
			return err
		}
		err = tx.Save(senderWallet).Error
		if err != nil {
			return err
		}

		return nil
	})
	return err
}

func (u *TransactionRepository) TransferList() ([]Transaction, error) {
	transactions := []Transaction{}
	err := u.db.Where("type = ?", "Credit").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
