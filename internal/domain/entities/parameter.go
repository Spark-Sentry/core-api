package entities

import (
	"gorm.io/gorm"
)

// Parameter represents a configuration or measurement parameter in the system.
// It has a many-to-one relationship with the Equipment entity.
type Parameter struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"` // Primary key, auto-incremented
	Name        string    `gorm:"size:255;not null"`        // Name of the parameter
	HostDevice  int       `gorm:"not null"`                 // Host device identifier (7 digits)
	Device      int       `gorm:"not null"`                 // Device identifier (7 digits)
	Log         int64     `gorm:"not null"`                 // Log number (10 digits)
	Point       string    `gorm:"size:255;not null"`        // Point reference (e.g., "AV-8084")
	Unit        string    `gorm:"size:50;not null"`         // Unit of measurement (e.g., "CFM")
	EquipmentID uint      `gorm:"not null"`                 // Foreign key linking to Equipment
	Equipment   Equipment `gorm:"foreignKey:EquipmentID"`   // Many-to-one relationship with Equipment
	gorm.Model            // GORM's built-in fields (ID, CreatedAt, UpdatedAt, DeletedAt)
}
