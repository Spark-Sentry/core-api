package entities

import "time"

type Contract struct {
	ID            uint      `gorm:"primaryKey"`
	AccountID     uint      `gorm:"not null"`
	StartDate     time.Time `gorm:"not null"`
	EndDate       time.Time `gorm:"not null"`
	ContractTerms string    `gorm:"type:text"` // Conditions et termes du contrat
	IsActive      bool      `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
