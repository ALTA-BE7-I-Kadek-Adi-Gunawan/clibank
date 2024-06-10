package cmd

import (
	"context"
	"fmt"

	"github.com/1layar/clibank/app/users"
	"github.com/1layar/clibank/platform"
	"github.com/AlecAivazis/survey/v2"
)

type CmdAddUser struct {
	Questions []*survey.Question
}

func (c *CmdAddUser) BuidQuestion() {
	if c.Questions == nil {
		c.Questions = []*survey.Question{
			{
				Name: "name",
				Prompt: &survey.Input{
					Message: "Enter your name: ",
				},
			},
			{
				Name: "phone",
				Prompt: &survey.Input{
					Message: "Enter Phone Number: ",
				},
			},
			{
				Name: "email",
				Prompt: &survey.Input{
					Message: "Enter Email Address: ",
				},
			},
			{
				Name: "pin",
				Prompt: &survey.Password{
					Message: "Enter new pin: ",
				},
			},
			{
				Name: "confirm_pin",
				Prompt: &survey.Password{
					Message: "Enter Pin Again: ",
				},
			},
		}
	}
}
func (c *CmdAddUser) Execute(ctx context.Context) error {
	service := ctx.Value(platform.UserServiceKey)
	user := &users.CreateUserDto{}
	c.BuidQuestion()
	survey.Ask(c.Questions, user)
	var userService users.UserService = service.(users.UserService)
	userDb, err := userService.CreateUser(*user)
	if err != nil {
		fmt.Printf("\nError: %v\n", err)
		return err
	}

	fmt.Printf("\nUser %v has been created\n", userDb.ID)
	return nil
}
