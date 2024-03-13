package entities

import "gorm.io/gorm"

type System struct {
	gorm.Model
	Name        string      `gorm:"size:255;not null"`
	Description string      `gorm:"size:255"`
	BuildingID  uint        `gorm:"not null"`
	Building    Building    `gorm:"foreignKey:BuildingID"`
	Equipments  []Equipment `gorm:"foreignKey:SystemID"`
}
