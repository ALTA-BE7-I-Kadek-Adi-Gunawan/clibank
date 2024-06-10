package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string   `gorm:"type:varchar(100);unique;not null" json:"email"`
	Pin         string   `gorm:"type:varchar(100);not null" json:"pin"`
	PhoneNumber string   `gorm:"type:varchar(12);uniqueIndex;not null" json:"phone_number"`
	Account     *Account `gorm:"foreignKey:PhoneNumber;references:PhoneNumber;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"account"`
}
