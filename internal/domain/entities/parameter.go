package entities

import "gorm.io/gorm"

type Parameter struct {
	Name        string `gorm:"size:255;not null"`
	HostDevice  int    `gorm:"type:int(7)"`
	Device      int    `gorm:"type:int(7)"`
	Log         int    `gorm:"type:int(10)"`
	Point       string `gorm:"size:255"`
	EquipmentID uint   `gorm:"not null"`
	Unit        string `gorm:"size:255"`
	gorm.Model
}
