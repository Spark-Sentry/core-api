package entities

import "gorm.io/gorm"

type Equipment struct {
	Tag         string      `gorm:"size:255;not null;unique"`
	Description string      `gorm:"size:255"`
	SystemID    uint        `gorm:"not null"`
	System      System      `gorm:"foreignKey:SystemID"`
	Parameters  []Parameter `gorm:"foreignKey:EquipmentID"`
	gorm.Model
}
