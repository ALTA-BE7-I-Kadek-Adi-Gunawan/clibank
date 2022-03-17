package topups

import "time"

type TopupHistory struct {
	Name      string
	BankName  string
	Status    string
	AccNo     string
	Amount    float64
	CreatedAt time.Time
}
