package transactions

import (
	"errors"

	"github.com/1layar/clibank/app/users"
)

type ITransactionService interface {
	Transfer(receiverPhone string, senderPhone string, ammount float64) error
	TransferList() ([]Transaction, error)
}

type TransactionService struct {
	repository  ITransactionRepository
	userService users.UserService
}

func (s *TransactionService) Init(transferRepo ITransactionRepository, userService users.UserService) {
	s.repository = transferRepo
	s.userService = userService
}

func (u *TransactionService) Transfer(receiverPhone string, senderPhone string, ammount float64) error {
	receiver, err := u.userService.GetUser(receiverPhone)
	if err != nil {
		return err
	}

	sender, err := u.userService.GetUser(senderPhone)
	if err != nil {
		return err
	}
	if sender.Account.Wallet.Balance < ammount && sender.Account.Wallet.Balance != 0 {
		return errors.New("balance is not sufficient to do transfer")
	}
	return u.repository.Transfer(receiver.Account, sender.Account, ammount)
}

func (u *TransactionService) TransferList() ([]Transaction, error) {
	return u.repository.TransferList()
}
