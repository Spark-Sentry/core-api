package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

type BuildingRepository struct {
	db *gorm.DB
}

func NewBuildingRepository(db *gorm.DB) *BuildingRepository {
	return &BuildingRepository{db: db}
}

// CreateBuilding creates a new Building with its associated Areas in the database.
func (r *BuildingRepository) CreateBuilding(building *entities.Building) error {
	// Use a transaction to ensure atomicity
	return r.db.Transaction(func(tx *gorm.DB) error {
		areas := building.Areas
		building.Areas = nil

		if err := tx.Create(building).Error; err != nil {
			return err
		}

		for i := range areas {
			areas[i].BuildingID = building.ID
			if err := tx.Create(&areas[i]).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

func (r *BuildingRepository) FindAllByAccountID(accountID uint) ([]entities.Building, error) {
	var buildings []entities.Building
	err := r.db.Where("account_id = ?", accountID).Preload("Areas").Find(&buildings).Error
	return buildings, err
}
