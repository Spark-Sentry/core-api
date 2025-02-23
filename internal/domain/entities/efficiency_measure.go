package entities

import "gorm.io/gorm"

// EfficiencyMeasure represents a measure (e.g. "HVAC Optimisation").
// CHANGES:
//   - Added fields: EfficiencyCost, MaintenanceCost,
//     Targets ([]Target) and Parameters ([]Parameter).
//
// This is a completely new entity.
type EfficiencyMeasure struct {
	ProjectID       uint        `gorm:"not null"`
	Name            string      `gorm:"size:255;not null"` // e.g. "HVAC Optimisation"
	EfficiencyCost  float64     `gorm:"not null"`
	MaintenanceCost float64     `gorm:"not null"`
	Targets         []Target    `gorm:"foreignKey:EfficiencyMeasureID"`
	Parameters      []Parameter `gorm:"foreignKey:EfficiencyMeasureID"`
	gorm.Model
}
