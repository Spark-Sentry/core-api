package entities

import "gorm.io/gorm"

// Category represents a building category.
// CHANGES:
// - Added a self-referencing relation "Parent" to allow nested categories.
// If a Category table already exists, it must be migrated.
type Category struct {
	Name     string    `gorm:"size:255;not null"` // e.g. "Retail store"
	ParentID *uint     // Nullable foreign key for parent category
	Parent   *Category `gorm:"foreignKey:ParentID"`
	gorm.Model
}
