package topups

import (
	"gorm.io/gorm"
)

type ITopupRepository interface {
	Topup(topupWallet *TopupWallet) error
	GetTopupOptions() ([]TopupOption, error)
	GetTopupOption(method string) (TopupOption, error)
	CreateTopupOption(topupOption *TopupOption) error
	GetTopUpHistory() ([]TopupHistory, error)
}

type TopupRepository struct {
	db *gorm.DB
}

func (rep *TopupRepository) Init(db *gorm.DB) {
	rep.db = db
}

func (repo *TopupRepository) Topup(topupWallet *TopupWallet) error {
	return repo.db.Create(topupWallet).Error
}

func (repo *TopupRepository) GetTopupOptions() ([]TopupOption, error) {
	topupOption := []TopupOption{}
	err := repo.db.Find(&topupOption).Error
	return topupOption, err
}

func (repo *TopupRepository) GetTopupOption(method string) (TopupOption, error) {
	topupOption := TopupOption{}
	err := repo.db.Where("code = ?", method).First(&topupOption).Error
	return topupOption, err
}

func (repo *TopupRepository) CreateTopupOption(topupOption *TopupOption) error {
	return repo.db.Create(topupOption).Error
}

func (repo *TopupRepository) GetTopUpHistory() ([]TopupHistory, error) {
	histories := []TopupHistory{}
	paymentJoin := "LEFT JOIN topup_options ON topup_options.id = topup_wallets.topup_option_id"
	walletJoin := "LEFT JOIN wallets ON wallets.id = topup_wallets.wallet_id"
	accountJoin := "LEFT JOIN accounts ON accounts.wallet_id = topup_wallets.wallet_id"
	historySelector := "accounts.name, topup_options.name as bank_name, topup_options.code, topup_options.acc_no, topup_options.status, topup_wallets.amount, topup_wallets.created_at"
	err := repo.db.Table("topup_wallets").Select(historySelector).Joins(paymentJoin).Joins(walletJoin).Joins(accountJoin).Scan(&histories).Error
	return histories, err
}
