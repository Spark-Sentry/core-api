package entities

import "gorm.io/gorm"

type Building struct {
	ID        uint   `gorm:"primaryKey"`            // Identifiant unique du bâtiment
	AccountID uint   `gorm:"not null"`              // ID du compte associé au bâtiment
	Name      string `gorm:"size:255;not null"`     // Nom du bâtiment
	Address   string `gorm:"size:255"`              // Adresse du bâtiment
	Group     string `gorm:"type:text"`             // Groupe ou catégorie du bâtiment
	Areas     []Area `gorm:"foreignKey:BuildingID"` // Les zones associées à ce bâtiment
	gorm.Model
}
