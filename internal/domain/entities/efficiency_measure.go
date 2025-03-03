package entities

import "gorm.io/gorm"

// EfficiencyMeasure represents a measure (e.g. "HVAC Optimisation").
type EfficiencyMeasure struct {
	ProjectID       uint        `gorm:"not null"`
	Name            string      `gorm:"size:255;not null"`
	EfficiencyCost  float64     `gorm:"not null"`
	MaintenanceCost float64     `gorm:"not null"`
	Targets         []Target    `gorm:"foreignKey:EfficiencyMeasureID"`
	Parameters      []Parameter `gorm:"many2many:parameter_efficiency_measures;"`
	gorm.Model
}
