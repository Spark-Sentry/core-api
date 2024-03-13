package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

type SystemRepository struct {
	db *gorm.DB
}

// NewSystemRepository creates a new instance of SystemRepository.
func NewSystemRepository(db *gorm.DB) *SystemRepository {
	return &SystemRepository{db: db}
}

// CreateSystem creates a new System in the database.
func (r *SystemRepository) CreateSystem(system *entities.System) error {
	return r.db.Create(system).Error
}

// FindSystemsByBuildingID retrieves all Systems associated with a given BuildingID.
func (r *SystemRepository) FindSystemsByBuildingID(buildingID uint) ([]entities.System, error) {
	var systems []entities.System
	err := r.db.Where("building_id = ?", buildingID).Find(&systems).Error
	return systems, err
}

func (r *SystemRepository) FindByBuildingID(buildingID uint) ([]entities.System, error) {
	var systems []entities.System
	err := r.db.Where("building_id = ?", buildingID).Find(&systems).Error
	return systems, err
}
