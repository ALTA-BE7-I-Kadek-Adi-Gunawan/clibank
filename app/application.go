package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/topups"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/transactions"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/users"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/wallets"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/cmd"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/controller"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/platform"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

const (
	cleanliness = "\033[H\033[2J"
	divider     = "==============================================================\n"
	spacing     = "|                                                            |\n"
)

type Application struct {
	choice          int8
	errorMsg        string
	display         string
	config          *platform.Configuration
	cmds            map[int8]Command
	ctx             context.Context
	database        *platform.Database
	usersController *controller.UserController
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
	transferRepo := &transactions.TransactionRepository{}
	userRepo.Init(db.DB)
	topupRepo.Init(db.DB)
	walletRepo.Init(db.DB)
	transferRepo.Init(db.DB)

	// init all services
	userService := &users.UserService{}
	topupService := &topups.TopupService{}
	walletService := &wallets.WalletService{}
	transactionService := &transactions.TransactionService{}
	userService.Init(userRepo)
	walletService.Init(walletRepo)
	topupService.Init(walletService, topupRepo)
	transactionService.Init(transferRepo, *userService)

	usersController := &controller.UserController{}
	usersController.Init(userService)
	a.usersController = usersController
	// seed data
	topupService.SeedOption()

	a.ctx = context.WithValue(a.ctx, platform.UserRepositoryKey, *userRepo)
	a.ctx = context.WithValue(a.ctx, platform.UserServiceKey, *userService)
	a.ctx = context.WithValue(a.ctx, platform.TopupRepositoryKey, *topupRepo)
	a.ctx = context.WithValue(a.ctx, platform.TopupServiceKey, *topupService)
	a.ctx = context.WithValue(a.ctx, platform.TransactionServiceKey, *transactionService)

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

func (a *Application) ClearTerminal() error {
	// find total \n in string
	if a.display == "" {
		return errors.New("Display is empty")
	}
	lines := strings.Count(a.display, "\n")
	if runtime.GOOS == "windows" {
		for i := 0; i < lines; i++ {
			fmt.Printf("\033[F\033[K")
		}
	} else if runtime.GOOS == "linux" {
		for i := 0; i < lines; i++ {
			fmt.Print("\033[1A")
		}
	} else {
		return errors.New("OS not supported")
	}

	return nil

}

func (a *Application) ShowMenu() error {
	var choice int8
	print("\nEnter your choice: \n")
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		a.errorMsg = "\nInvalid input, only accept number!\n"
		return err
	}

	if _, ok := a.cmds[choice]; ok {
		a.choice = choice
		return nil
	}

	if choice == 0 {
		print(a.ThankYou())
		a.choice = choice
		return errors.New("Exit")
	}

	a.errorMsg = "\nInvalid choice, please try again!\n"
	return errors.New("invalid choice")
}

func (a *Application) Update() string {
	output := divider
	output += a.ShowHeader()
	output += spacing
	if a.GetChoice() == 0 {
		output += a.ThankYou()
	} else {
		output += a.showInfo()
	}
	output += spacing
	output += divider
	if a.errorMsg != "" {
		output += a.errorMsg
		a.errorMsg = ""
	}

	return output
}

func (a Application) ThankYou() string {
	message := "Terima Kasih Telah bertansaki dengan kami!\n"
	return message
}

func (a *Application) Run(s ...string) error {
	runMode := s[0]

	switch runMode {
	case "server":
		a.RunServer()
	case "cli":
		a.RunCLI()
	case "help":
		fmt.Println("Usage:")
		fmt.Println("\tclibank server")
		fmt.Println("\tclibank cli")
	default:
		return errors.New("invalid run mode")
	}
	return nil
}

func (a *Application) RunServer() {

	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/api/v1/users", a.usersController.GetUsers)
	e.GET("/api/v1/users/:id", a.usersController.GetUser)
	e.PUT("/api/v1/users/:id", a.usersController.UpdateUser)
	e.POST("/api/v1/users", a.usersController.CreateUser)
	e.DELETE("/api/v1/users/:id", a.usersController.DeleteUser)

	e.Logger.Fatal(e.Start(":8000"))
}

func (a *Application) RunCLI() {
	a.display = a.Update()
	fmt.Print(a.display)
	if val, ok := a.cmds[a.choice]; ok {
		val.Execute(a.ctx)
	} else {
		if a.choice > 0 {
			a.errorMsg = "Invalid choice!"
		}
	}
}

type Command interface {
	Execute(ctx context.Context) error
}
