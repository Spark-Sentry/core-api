package entities

import "gorm.io/gorm"

// Coefficient represents a coefficient value used in regressions.
// CHANGES:
// - Added field "Value" (float) and an optional link to an IndependantVariable.
// - Added "RegressionID" to associate with a Regression.
type Coefficient struct {
	Value                 float64              `gorm:"not null"` // e.g. 363.4
	IndependantVariableID *uint                // Optional relation
	IndependantVariable   *IndependantVariable `gorm:"foreignKey:IndependantVariableID"`
	RegressionID          uint                 `gorm:"not null"`
	gorm.Model
}
