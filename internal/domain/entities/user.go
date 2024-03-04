package entities

import (
	"gorm.io/gorm"
)

type User struct {
	Name      string
	Email     string `gorm:"unique"`
	AccountID uint
	Password  string
	Account   Account
	gorm.Model
}
