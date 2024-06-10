package responses

import (
	"github.com/1layar/clibank/app/users"
	"github.com/1layar/clibank/app/wallets"
)

type ApiResponse struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Error      string      `json:"error"`
	Data       interface{} `json:"data"`
}

type UsersApiResponse struct {
	Message    string         `json:"message"`
	StatusCode int            `json:"status_code"`
	Data       []UserResponse `json:"data"`
}

type UserApiResponse struct {
	Message    string       `json:"message"`
	StatusCode int          `json:"status_code"`
	Data       UserResponse `json:"data"`
}

type WalletResponse struct {
	// ID        uint    `json:"id"`
	Balance   float64 `json:"balance"`
	Status    string  `json:"status"`
	Currency  string  `json:"currency"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func (wallet *WalletResponse) FromModel(model *wallets.Wallet) {
	// wallet.ID = model.ID
	wallet.Balance = model.Balance
	wallet.Status = model.Status
	wallet.Currency = model.Currency
	wallet.CreatedAt = model.CreatedAt.String()
	wallet.UpdatedAt = model.UpdatedAt.String()
}

type AccountResponse struct {
	// ID          uint           `json:"id"`
	Name string `json:"name"`
	// PhoneNumber string         `json:"phone_number"`
	Wallet    WalletResponse `json:"wallet"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}

func (account *AccountResponse) FromModel(userModel users.Account) {
	// account.ID = userModel.ID
	account.Name = userModel.Name
	// account.PhoneNumber = userModel.PhoneNumber
	account.CreatedAt = userModel.CreatedAt.String()
	account.UpdatedAt = userModel.UpdatedAt.String()
	if userModel.Wallet != nil {
		account.Wallet = WalletResponse{}
		account.Wallet.FromModel(userModel.Wallet)
	}
}

type UserResponse struct {
	ID          uint            `json:"id"`
	Email       string          `json:"email"`
	PhoneNumber string          `json:"phone_number"`
	Account     AccountResponse `json:"account,omitempty"`
	CreatedAt   string          `json:"created_at"`
	UpdatedAt   string          `json:"updated_at"`
}

func (user *UserResponse) FromModel(userModel users.User) {
	user.ID = userModel.ID
	user.Email = userModel.Email
	user.PhoneNumber = userModel.PhoneNumber
	user.CreatedAt = userModel.CreatedAt.String()
	user.UpdatedAt = userModel.UpdatedAt.String()
	if userModel.Account != nil {
		user.Account = AccountResponse{}
		user.Account.FromModel(*userModel.Account)
	}
}
