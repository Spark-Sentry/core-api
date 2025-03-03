package entities

import "gorm.io/gorm"

// Bill represents a billing record for a building.
type Bill struct {
	gorm.Model
	BuildingID uint    `gorm:"not null"`         // Foreign key linking to Building
	MeterID    uint    `gorm:"not null"`         // Associated meter ID
	Start      string  `gorm:"size:50;not null"` // Start datetime (e.g., RFC3339 format)
	Stop       string  `gorm:"size:50;not null"` // End datetime
	Quantity   float64 `gorm:"not null"`         // Quantity measured
	Cost       float64 `gorm:"not null"`         // Cost associated with the bill
}
