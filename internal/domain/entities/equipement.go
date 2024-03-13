package entities

import "gorm.io/gorm"

type Equipment struct {
	gorm.Model
	Tag         string `gorm:"size:255;not null;unique"`
	Description string `gorm:"size:255"`
	SystemID    uint   `gorm:"not null"`
	System      System `gorm:"foreignKey:SystemID"`
}
