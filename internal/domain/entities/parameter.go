package entities

import "gorm.io/gorm"

// Parameter represents a measurement parameter.
// CHANGES:
// - "EquipmentID" is optional (the parameter may not be linked to an equipment).
// - Added field "BmsID" (required) with relation to Bms.
// - Added fields "IDInBms", "PointType" and "Unit" as strings.
// - Optional link to an EfficiencyMeasure via EfficiencyMeasureID.
type Parameter struct {
	Name                string     `gorm:"size:255;not null"`
	EquipmentID         *uint      // Optional
	Equipment           *Equipment `gorm:"foreignKey:EquipmentID"`
	BmsID               uint       `gorm:"not null"`
	Bms                 Bms        `gorm:"foreignKey:BmsID"`
	IDInBms             string     `gorm:"size:255;not null"` // e.g. "0027400.0027401.0000000073.AV-8084"
	PointType           string     `gorm:"size:50;not null"`  // e.g. "AV"
	Unit                string     `gorm:"size:50;not null"`  // e.g. "°C"
	EfficiencyMeasureID *uint      // Optional link
	gorm.Model
}
