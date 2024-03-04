package entities

import "gorm.io/gorm"

type Building struct {
	ID          uint   `gorm:"primaryKey"`
	AccountID   uint   `gorm:"not null"`
	Name        string `gorm:"size:255;not null"`
	Address     string `gorm:"size:255"`
	Description string `gorm:"type:text"`
	Category    string // Catégorisation du bâtiment
	gorm.Model
}
