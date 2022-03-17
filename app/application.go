package app

import (
	"context"
	"fmt"

	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/topups"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/users"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/wallets"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/cmd"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/platform"
)

const (
	cleanliness = "\033[H\033[2J"
	divider     = "==============================================================\n"
	spacing     = "|                                                            |\n"
)

type Application struct {
	choice   int8
	config   *platform.Configuration
	cmds     map[int8]Command
	ctx      context.Context
	database *platform.Database
}

func (a *Application) GetChoice() int8 {
	return a.choice
}

func (a *Application) SetChoice(choice int8) {
	a.choice = choice
}

func (a *Application) Init(db *platform.Database, c *platform.Configuration) {
	a.database = db
	a.choice = -1
	a.config = c
	a.ctx = context.Background()

	// init all repositores
	userRepo := &users.UserRepository{}
	topupRepo := &topups.TopupRepository{}
	walletRepo := &wallets.WalletRepository{}
	userRepo.Init(db.DB)
	topupRepo.Init(db.DB)
	walletRepo.Init(db.DB)

	// init all services
	userService := &users.UserService{}
	topupService := &topups.TopupService{}
	walletService := &wallets.WalletService{}
	userService.Init(userRepo)
	walletService.Init(walletRepo)
	topupService.Init(walletService, topupRepo)

	// seed data
	topupService.SeedOption()

	a.ctx = context.WithValue(a.ctx, platform.UserRepositoryKey, *userRepo)
	a.ctx = context.WithValue(a.ctx, platform.UserServiceKey, *userService)
	a.ctx = context.WithValue(a.ctx, platform.TopupRepositoryKey, *topupRepo)
	a.ctx = context.WithValue(a.ctx, platform.TopupServiceKey, *topupService)

	a.cmds = map[int8]Command{
		1: &cmd.CmdAddUser{},
		2: &cmd.CmdUpdateUser{},
		3: &cmd.CmdDeleteUser{},
		4: &cmd.CmdGetUser{},
		5: &cmd.CmdAccoutnTopUp{},
		6: cmd.CmdTransferBalance{},
		7: cmd.CmdHistoryTopUp{},
		8: cmd.CmdHistoryTransaction{},
	}
}
func (a Application) ShowHeader() string {
	header := "|   $$$$$$\\  $$\\ $$\\ $$$$$$$\\                      $$\\       |\n"
	header += "|  $$  __$$\\ $$ |\\__|$$  __$$\\                     $$ |      |\n"
	header += "|  $$ /  \\__|$$ |$$\\ $$ |  $$ | $$$$$$\\  $$$$$$$\\  $$ |  $$\\ |\n"
	header += "|  $$ |      $$ |$$ |$$$$$$$\\ | \\____$$\\ $$  __$$\\ $$ | $$  ||\n"
	header += "|  $$ |      $$ |$$ |$$  __$$\\  $$$$$$$ |$$ |  $$ |$$$$$$  / |\n"
	header += "|  $$ |  $$\\ $$ |$$ |$$ |  $$ |$$  __$$ |$$ |  $$ |$$  _$$<  |\n"
	header += "|  \\$$$$$$  |$$ |$$ |$$$$$$$  |\\$$$$$$$ |$$ |  $$ |$$ | \\$$\\ |\n"
	header += "|  \\______/ \\__|\\__|\\_______/  \\_______|\\__|  \\__|\\__|  \\__| |\n"
	header += "|                                                            |\n"

	return header
}
func (a Application) showInfo() string {
	info := "|  1. Add User                                               |\n"
	info += "|  2. Update User                                            |\n"
	info += "|  3. Delete User                                            |\n"
	info += "|  4. Get User                                               |\n"
	info += "|  5. Top Up Wallet Balance                                  |\n"
	info += "|  6. Transfer To Other User                                 |\n"
	info += "|  7. History Top Up                                         |\n"
	info += "|  8. History Transfer                                       |\n"
	info += "|  0. Exit                                                   |\n"
	return info
}
func (a *Application) ShowMenu() {
	var choice int8
	print("Enter your choice: ")
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		println("Invalid input, only accept number from 1 to 0!")
	}
	a.choice = choice
}

func (a *Application) Update() string {
	output := divider
	output += a.ShowHeader()
	output += spacing
	output += a.showInfo()
	output += spacing
	output += divider

	return output
}

func (a *Application) Run() {
	fmt.Print(a.Update())
	if a.choice > 0 {
		a.cmds[a.choice].Execute(a.ctx)
	} else {
		if a.choice == 0 {
			println("Terima Kasih Telah bertansaki dengan kami!")
		}
	}
}

type Command interface {
	Execute(ctx context.Context) error
}
