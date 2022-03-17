package cmd

import (
	"context"
	"fmt"

	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/topups"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/platform"
)

type CmdHistoryTopUp struct {
}

func (c CmdHistoryTopUp) Execute(ctx context.Context) error {
	service := ctx.Value(platform.TopupServiceKey)
	var topupService topups.TopupService = service.(topups.TopupService)
	history, err := topupService.GetTopUpHistory()

	if err != nil {
		return err
	}

	for _, history := range history {
		strDate := history.CreatedAt.Format("2006-01-02 15:04")
		fmt.Printf("name: %v \nbank: %v \nbank_num: %v \namount: %.2f \nstatus: %v \ndate: %v \n", history.Name, history.BankName, history.AccNo, history.Amount, history.Status, strDate)
	}

	return nil
}
