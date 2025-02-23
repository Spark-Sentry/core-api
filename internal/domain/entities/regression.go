package entities

import "gorm.io/gorm"

// Regression represents a regression model.
// CHANGES:
//   - Added fields: Name (e.g. "Monthly electricity consumption"), Unit,
//     Coefficients ([]Coefficient) and a many-to-many relation with Meters.
type Regression struct {
	ProjectID    uint          `gorm:"not null"`          // Associated project
	Name         string        `gorm:"size:255;not null"` // e.g. "Monthly electricity consumption"
	Unit         string        `gorm:"size:50;not null"`  // e.g. "kWh"
	Coefficients []Coefficient `gorm:"foreignKey:RegressionID"`
	Meters       []Meter       `gorm:"many2many:regression_meters;"`
	gorm.Model
}
