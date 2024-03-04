package entities

import (
	"gorm.io/gorm"
)

type Account struct {
	Name  string
	Users []User
	gorm.Model
}
