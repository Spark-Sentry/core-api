package entities

import "gorm.io/gorm"

// Parameter represents a measurement parameter.
type Parameter struct {
	gorm.Model
	Name        string     `gorm:"size:255;not null"`
	EquipmentID *uint      // Optional
	Equipment   *Equipment `gorm:"foreignKey:EquipmentID"`
	BmsID       uint       `gorm:"not null"`
	Bms         Bms        `gorm:"foreignKey:BmsID"`
	IDInBms     string     `gorm:"size:255;not null"`
	PointType   string     `gorm:"size:50;not null"`
	Unit        string     `gorm:"size:50;not null"`
	// Many-to-many relation with EfficiencyMeasure:
	EfficiencyMeasures []EfficiencyMeasure `gorm:"many2many:parameter_efficiency_measures;"`
}
