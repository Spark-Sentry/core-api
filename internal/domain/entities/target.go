package entities

import "gorm.io/gorm"

// Target represents a target value (e.g., "Simple return on investment period").
// Fields:
// - Name: Name of the target (e.g., "Simple return on investment period")
// - Value: Numeric value (e.g., 3.4)
// - Type: Type of target (e.g., "PBP")
// It can be associated with either a Project or an EfficiencyMeasure.
type Target struct {
	Name                string  `gorm:"size:255;not null"`
	Value               float64 `gorm:"not null"`
	Type                string  `gorm:"size:50;not null"`
	ProjectID           *uint   // Optional: linked to a Project
	EfficiencyMeasureID *uint   // Optional: linked to an EfficiencyMeasure
	gorm.Model
}
