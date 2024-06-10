package cmd

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"

	"github.com/1layar/clibank/app/users"
	"github.com/1layar/clibank/platform"
)

type CmdUpdateUser struct {
	Phone      string
	Pin        string
	UpdatePin  string `survey:"update_pin"`
	UpdateName string `survey:"update_name"`
	Questions  []*survey.Question
}

func (cmd *CmdUpdateUser) BuidQuestion() {
	if cmd.Questions == nil {
		cmd.Questions = []*survey.Question{
			{
				Name: "phone",
				Prompt: &survey.Input{
					Message: "Enter Phone Number: ",
				},
			},
			{
				Name: "pin",
				Prompt: &survey.Password{
					Message: "Enter your pin: ",
				},
			},
			{
				Name: "update_name",
				Prompt: &survey.Input{
					Message: "Enter your new name: ",
				},
			},
			{
				Name: "update_pin",
				Prompt: &survey.Password{
					Message: "Enter new pin: ",
				},
			},
		}
	}
}

func (c *CmdUpdateUser) Execute(ctx context.Context) error {
	service := ctx.Value(platform.UserServiceKey)
	user := &users.UpdateUserDto{}
	c.BuidQuestion()
	survey.Ask(c.Questions, c)
	user.Name = c.UpdateName
	user.Pin = c.UpdatePin
	var userService users.UserService = service.(users.UserService)
	userDb, err := userService.UpdateUser(c.Phone, c.Pin, *user)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	fmt.Printf("User %v has been updated\n", userDb.ID)
	return nil
}
