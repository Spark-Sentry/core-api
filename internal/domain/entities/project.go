package entities

import "gorm.io/gorm"

// Project represents a project linked to a building.
// CHANGES:
//   - Added fields: ContractorID, Contractor relation, ImplementationDate (string),
//     EfficiencyCost, MaintenanceCost, Targets ([]Target),
//     EfficiencyMeasures ([]EfficiencyMeasure) and Subsidies ([]Subsidy).
//
// This is a completely new entity.
type Project struct {
	BuildingID         uint                `gorm:"not null"`
	Name               string              `gorm:"size:255;not null"`
	ContractorID       uint                `gorm:"not null"`
	Contractor         Contractor          `gorm:"foreignKey:ContractorID"`
	ImplementationDate string              `gorm:"size:50;not null"`
	EfficiencyCost     float64             `gorm:"not null"`
	MaintenanceCost    float64             `gorm:"not null"`
	Targets            []Target            `gorm:"foreignKey:ProjectID"`
	EfficiencyMeasures []EfficiencyMeasure `gorm:"foreignKey:ProjectID"`
	Subsidies          []Subsidy           `gorm:"foreignKey:ProjectID"`
	gorm.Model
}
