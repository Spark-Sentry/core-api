package entities

import "gorm.io/gorm"

// Meter represents a meter installed in a building.
// CHANGES:
// - Added fields: Supplier, Energy, Name, and a unique "MeterID" (string).
// - Linked to a Building.
type Meter struct {
	Supplier   string   `gorm:"size:255;not null"` // e.g. "Hydro-Quebec"
	BuildingID uint     `gorm:"not null"`
	Building   Building `gorm:"foreignKey:BuildingID"`
	Energy     string   `gorm:"size:50;not null"`         // e.g. "Electricity"
	Name       string   `gorm:"size:255;not null"`        // e.g. "Main"
	MeterID    string   `gorm:"size:255;not null;unique"` // e.g. "644090824"
	gorm.Model
}
