package entities

import "gorm.io/gorm"

// Contractor represents a contractor with contact details.
// This is a completely new entity.
type Contractor struct {
	Name  string `gorm:"size:255;not null"`
	Phone string `gorm:"size:50"`
	Email string `gorm:"size:255"`
	gorm.Model
}
