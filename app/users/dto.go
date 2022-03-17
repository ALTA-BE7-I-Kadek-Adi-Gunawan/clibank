package users

type CreateUserDto struct {
	Email      string `json:"email" survey:"email"`
	Phone      string `json:"phone" survey:"phone"`
	Name       string `json:"name" survey:"name"`
	Pin        string `json:"pin" survey:"pin"`
	ConfirmPin string `json:"confirm_pin" survey:"confirm_pin"`
}

type UpdateUserDto struct {
	Name string `json:"name"`
	Pin  string `json:"pin"`
}
