package entities

import "gorm.io/gorm"

// Building represents a building in the system.
type Building struct {
	gorm.Model
	AccountID  uint      `gorm:"not null"`              // Associated account ID
	Name       string    `gorm:"size:255;not null"`     // Building name
	Address    string    `gorm:"size:255"`              // Building address
	Group      string    `gorm:"type:text"`             // Group or category of the building
	Bills      []Bill    `gorm:"foreignKey:BuildingID"` // Associated bills
	Projects   []Project `gorm:"foreignKey:BuildingID"` // Associated projects
	CategoryID *uint     // Optional: linked category ID
	Category   Category  // Associated category (if any)
}
