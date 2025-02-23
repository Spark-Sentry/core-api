package entities

import "gorm.io/gorm"

// Equipment represents a piece of equipment (e.g. "Make-up air unit").
// CHANGES:
// - Removed the SystemID and Description fields.
// - Added a one-to-many relationship with Parameter.
type Equipment struct {
	gorm.Model
	Name       string      `gorm:"size:255;not null"`        // e.g., "Make-up air unit"
	Tag        string      `gorm:"size:255;not null;unique"` // e.g., "MUA-2"
	Parameters []Parameter `gorm:"foreignKey:EquipmentID"`   // Associated parameters
}
