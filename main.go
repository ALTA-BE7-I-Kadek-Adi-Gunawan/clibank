package main

import (
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/platform"
)

var (
	cliBank *app.Application
)

func init() {
	config := new(platform.Configuration)
	db := new(platform.Database)

	defer func() {
		if r := recover(); r != nil {
			println("Recovered from panic:", r)
		}
	}()

	config.InitConfiguration()
	db.Connect(config.DatabaseUrl)
	db.Migrate()
	cliBank = new(app.Application)
	cliBank.Init(db, config)
}

func main() {
	for cliBank.GetChoice() != 0 {
		// cliBank.ClearTerminal()
		cliBank.Run()
		cliBank.ShowMenu()
	}
}
