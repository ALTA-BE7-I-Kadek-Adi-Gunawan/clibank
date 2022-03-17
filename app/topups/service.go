package topups

import (
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/wallets"
	"github.com/google/uuid"
)

type ITopupService interface {
	Topup(option TopupWalletDTO) error
	CreateTopupOption(option AddTopupOptionDTO) error
	GetTopupOptions() ([]TopupOption, error)
	GetTopUpHistory() ([]TopupHistory, error)
}

type TopupService struct {
	walletServ wallets.IWalletService
	topupRepo  ITopupRepository
}

func (s *TopupService) Init(walletServ wallets.IWalletService, topupRepo ITopupRepository) {
	s.walletServ = walletServ
	s.topupRepo = topupRepo
}

func (s *TopupService) SeedOption() {
	s.CreateTopupOption(AddTopupOptionDTO{
		Name:  "Mandiri",
		Code:  "MANDIRI",
		AccNo: "1231321312",
	})
	s.CreateTopupOption(AddTopupOptionDTO{
		Name:  "BCA",
		Code:  "BCA",
		AccNo: "123123321",
	})
	s.CreateTopupOption(AddTopupOptionDTO{
		Name:  "BRI",
		Code:  "BRI",
		AccNo: "1231323",
	})
}

func (s TopupService) generateSerialNo() string {
	// panic anticipe
	return "CLIB-" + uuid.New().String()
}

func (s *TopupService) Topup(option TopupWalletDTO) (TopupWallet, error) {
	wallet, err := s.walletServ.GetWalletByPhoneNumber(option.PhoneNumber)
	if err != nil {
		return TopupWallet{}, err
	}

	topupOption, err := s.topupRepo.GetTopupOption(option.Method)
	if err != nil {
		return TopupWallet{}, err
	}

	topupWallet := &TopupWallet{
		Amount:        option.Amount,
		SerialNo:      s.generateSerialNo(),
		WalletID:      wallet.ID,
		TopupOptionID: topupOption.ID,
	}
	err = s.topupRepo.Topup(topupWallet)
	if err != nil {
		return TopupWallet{}, err
	}

	err = s.walletServ.UpdateWallet(wallets.UpdateWalletDto{
		ID:       wallet.ID,
		Currency: wallet.Currency,
		Balance:  wallet.Balance + option.Amount,
	})

	return *topupWallet, err
}

func (s *TopupService) CreateTopupOption(option AddTopupOptionDTO) error {
	topupOption := &TopupOption{
		Name:  option.Name,
		Code:  option.Code,
		AccNo: option.AccNo,
	}
	return s.topupRepo.CreateTopupOption(topupOption)
}

func (s *TopupService) GetTopupOptions() ([]TopupOption, error) {
	return s.topupRepo.GetTopupOptions()
}

func (s *TopupService) GetTopUpHistory() ([]TopupHistory, error) {
	return s.topupRepo.GetTopUpHistory()
}
