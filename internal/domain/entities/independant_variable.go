package entities

import "gorm.io/gorm"

// IndependantVariable represents an independent variable used in regressions.
// The Coefficient field is defined as a pointer to allow nil values
// when no coefficient is associated.
type IndependantVariable struct {
	gorm.Model
	Name        string       `gorm:"size:255;not null"`                              // e.g., "DJC 18°C MTL TRUDEAU"
	DateStart   string       `gorm:"size:50;not null"`                               // e.g., "2025-01-01T01:00:00Z"
	DateStop    string       `gorm:"size:50;not null"`                               // e.g., "2025-02-01T00:00:00Z"
	Value       float64      `gorm:"not null"`                                       // e.g., 439
	Source      string       `gorm:"size:50;not null"`                               // "input" or "computed"
	Coefficient *Coefficient `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Optional association
}
