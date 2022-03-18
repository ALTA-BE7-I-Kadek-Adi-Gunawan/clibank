package cmd

import (
	"context"
	"fmt"

	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/users"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/platform"
	"github.com/AlecAivazis/survey/v2"
)

type CmdGetUser struct {
	Phone     string `survey:"phone"`
	Questions []*survey.Question
}

func (c *CmdGetUser) BuildQuestion() {
	if c.Questions == nil {
		c.Questions = []*survey.Question{
			{
				Name: "phone",
				Prompt: &survey.Input{
					Message: "Enter Phone Number: ",
				},
			},
		}
	}
}

func (c *CmdGetUser) Execute(ctx context.Context) error {
	service := ctx.Value(platform.UserServiceKey)
	var userService users.UserService = service.(users.UserService)
	c.BuildQuestion()
	survey.Ask(c.Questions, c)
	defer func() {
		if r := recover(); r != nil {
			fmt.Print("\nUser Input Error\n", r)
		}
	}()
	userDb, err := userService.GetUser(c.Phone)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	fmt.Printf("User %v phone %v balance Rp.%.2f", userDb.Email, userDb.Account.PhoneNumber, userDb.Account.Wallet.Balance)

	return nil
}
