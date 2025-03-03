package entities

import "gorm.io/gorm"

// Bms represents a Building Management System (e.g. "Compass").
// CHANGES:
// - Added fields: Type and IDPattern (e.g. "Compass" and "{hostDevice}.{device}.{log}.{point}").
// This is a completely new entity.
type Bms struct {
	Type      string `gorm:"size:255;not null"` // e.g. "Compass"
	IDPattern string `gorm:"size:255;not null"` // e.g. "{hostDevice}.{device}.{log}.{point}"
	gorm.Model
}
