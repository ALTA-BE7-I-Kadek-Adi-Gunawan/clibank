package cmd

import (
	"context"
	"fmt"

	"github.com/1layar/clibank/app/topups"
	"github.com/1layar/clibank/platform"
	"github.com/AlecAivazis/survey/v2"
)

type CmdAccoutnTopUp struct {
	TopupMethod string  `survey:"topup_method"`
	PhoneNumber string  `survey:"phone_number"`
	Amount      float64 `survey:"amount"`
	Questions   []*survey.Question
}

func (cmd *CmdAccoutnTopUp) BuildQuestion(args ...string) {
	if cmd.Questions == nil {
		cmd.Questions = []*survey.Question{
			{
				Name: "phone_number",
				Prompt: &survey.Input{
					Message: "Enter Phone Number: ",
				},
			},
			{
				Name: "amount",
				Prompt: &survey.Input{
					Message: "Enter Amount: ",
				},
			},
			{
				Name: "topup_method",
				Prompt: &survey.Select{
					Message: "Select Topup Method: ",
					Options: args,
				},
			},
		}
	}
}

func (c CmdAccoutnTopUp) Execute(ctx context.Context) error {
	service := ctx.Value(platform.TopupServiceKey)
	var topupService topups.TopupService = service.(topups.TopupService)
	topupWallet := &topups.TopupWalletDTO{}
	options, err := topupService.GetTopupOptions()

	if err != nil {
		return err
	}
	choiceMap := make(map[string]string)
	choices := []string{}
	for _, option := range options {
		choices = append(choices, option.Name)
		choiceMap[option.Name] = option.Code
	}
	c.BuildQuestion(choices...)
	survey.Ask(c.Questions, &c)
	topupWallet.PhoneNumber = c.PhoneNumber
	topupWallet.Amount = c.Amount
	topupWallet.Method = choiceMap[c.TopupMethod]

	trans, err := topupService.Topup(*topupWallet)
	if err != nil {
		fmt.Println("failed to topup to wallet \t")
		return err
	}
	fmt.Printf("succesfully topup %v\n", trans.SerialNo)

	return nil
}
