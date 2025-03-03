package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

type BuildingRepository interface {
	CreateBuilding(building *entities.Building) error
	FindAllByAccountID(accountID uint) ([]entities.Building, error)
	UpdateBuilding(buildingID uint, updatedBuilding entities.Building) error
}

type buildingRepository struct {
	db *gorm.DB
}

func NewBuildingRepository(db *gorm.DB) BuildingRepository {
	return &buildingRepository{db: db}
}

func (r *buildingRepository) CreateBuilding(building *entities.Building) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(building).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *buildingRepository) FindAllByAccountID(accountID uint) ([]entities.Building, error) {
	var buildings []entities.Building
	err := r.db.Where("account_id = ?", accountID).
		Preload("Projects").
		Preload("Category").
		Preload("Bills").
		Find(&buildings).Error
	return buildings, err
}

func (r *buildingRepository) UpdateBuilding(buildingID uint, updatedBuilding entities.Building) error {
	return r.db.Model(&entities.Building{}).Where("id = ?", buildingID).Updates(map[string]interface{}{
		"name":        updatedBuilding.Name,
		"address":     updatedBuilding.Address,
		"category_id": updatedBuilding.CategoryID,
	}).Error
}
