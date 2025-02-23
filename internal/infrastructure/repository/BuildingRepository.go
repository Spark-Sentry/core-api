package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// BuildingRepository handles database operations for Building.
type BuildingRepository interface {
	CreateBuilding(building *entities.Building) error
	FindAllByAccountID(accountID uint) ([]entities.Building, error)
}

type buildingRepository struct {
	db *gorm.DB
}

// NewBuildingRepository creates a new instance of BuildingRepository.
func NewBuildingRepository(db *gorm.DB) BuildingRepository {
	return &buildingRepository{db: db}
}

// CreateBuilding creates a new Building record in the database.
// CHANGES:
// - The new Building entity does not include Areas.
// - A transaction is used to ensure atomicity.
// NOTE: The existing Building table must be migrated to include the new fields.
func (r *buildingRepository) CreateBuilding(building *entities.Building) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(building).Error; err != nil {
			return err
		}
		return nil
	})
}

// FindAllByAccountID retrieves all buildings associated with the given account ID.
// It preloads associated Projects, Category, and Bills.
func (r *buildingRepository) FindAllByAccountID(accountID uint) ([]entities.Building, error) {
	var buildings []entities.Building
	err := r.db.Where("account_id = ?", accountID).
		Preload("Projects").
		Preload("Category").
		Preload("Bills").
		Find(&buildings).Error
	return buildings, err
}
