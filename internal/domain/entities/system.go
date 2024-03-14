package entities

import "gorm.io/gorm"

type System struct {
	gorm.Model
	Name        string      `gorm:"size:255;not null"`
	Description string      `gorm:"size:255"`
	AreaID      uint        `gorm:"not null"`
	Area        Area        `gorm:"foreignKey:AreaID"`
	Equipments  []Equipment `gorm:"foreignKey:SystemID"`
}
