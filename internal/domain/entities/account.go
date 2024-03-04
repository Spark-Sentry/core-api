package entities

import (
	"gorm.io/gorm"
)

type Account struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:255;not null"`
	ContactEmail string `gorm:"size:255"`
	ContactPhone string `gorm:"size:20"`
	Plan         string `gorm:"size:50"` // Ex: Basic, Premium
	Users        []User
	Buildings    []Building
	gorm.Model
}
