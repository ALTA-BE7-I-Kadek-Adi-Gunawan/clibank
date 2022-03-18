package platform

import (
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/topups"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/transactions"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/users"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/wallets"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func (db *Database) Connect(url interface{}) {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	gormDb, err := gorm.Open(mysql.Open(url.(string)), &gorm.Config{
		// Hide sql logs
		//Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	db.DB = gormDb
}

func (db Database) Migrate() {
	db.AutoMigrate(&wallets.Wallet{})
	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&users.Account{})
	db.AutoMigrate(&topups.TopupWallet{})
	db.AutoMigrate(&topups.TopupOption{})
	db.AutoMigrate(&transactions.Transaction{})
}
