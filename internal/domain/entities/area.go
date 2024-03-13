package entities

import "gorm.io/gorm"

// Area represents a zone or space within a Building.
type Area struct {
	ID          uint   `gorm:"primaryKey"`        // Unique identifier for the Area
	BuildingID  uint   `gorm:"not null;index;"`   // Foreign key referencing Building
	Name        string `gorm:"size:255;not null"` // Name of the Area
	Description string `gorm:"size:255"`          // Optional description of the Area
	gorm.Model         // Includes standard Gorm fields (CreatedAt, UpdatedAt, DeletedAt)
}
