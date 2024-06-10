package cmd

import (
	"context"
	"fmt"

	"github.com/1layar/clibank/app/transactions"
	"github.com/1layar/clibank/platform"
)

type CmdTransferBalance struct {
	SenderPhone   string
	ReceiverPhone string
	Ammount       float64
}

func (c CmdTransferBalance) Execute(ctx context.Context) error {
	service := ctx.Value(platform.TransactionServiceKey)
	fmt.Println("Input sender phone no. : ")
	fmt.Scan(&c.SenderPhone)
	fmt.Println("Input receiver phone no. : ")
	fmt.Scan(&c.ReceiverPhone)
	fmt.Println("Input Ammount of Transfer ")
	fmt.Scan(&c.Ammount)
	var transactionService transactions.TransactionService = service.(transactions.TransactionService)
	err := transactionService.Transfer(c.ReceiverPhone, c.SenderPhone, c.Ammount)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
