package entities

import "gorm.io/gorm"

// Subsidy represents a subsidy record that can be linked to either a Project or an EfficiencyMeasure.
type Subsidy struct {
	Organisation        string  `gorm:"size:255;not null"` // e.g., "TEQ"
	DateObtained        string  `gorm:"size:50;not null"`  // Date of obtaining the subsidy
	Amount              float64 `gorm:"not null"`
	ProjectID           *uint   // Optional: if linked to a Project
	EfficiencyMeasureID *uint   // Optional: if linked to an EfficiencyMeasure
	gorm.Model
}
