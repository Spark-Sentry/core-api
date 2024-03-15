package repository

import (
	"core-api/internal/app/dto"
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

func (r *SystemRepository) FindByAreaID(areaID uint) ([]entities.System, error) {
	var systems []entities.System
	err := r.db.Where("area_id = ?", areaID).Preload("Area").Find(&systems).Error
	if err != nil {
		return nil, err
	}
	return systems, nil
}

// CreateEquipment creates equipment
func (r *SystemRepository) CreateEquipment(equipment *entities.Equipment) error {
	return r.db.Create(equipment).Error
}

func (r *SystemRepository) UpdateSystem(systemID uint, updateDTO dto.SystemUpdateDTO) error {
	return r.db.Model(&entities.System{}).Where("id = ?", systemID).Updates(entities.System{Name: updateDTO.Name, Description: updateDTO.Description}).Error
}

func (r *SystemRepository) DeleteSystem(systemID uint) error {
	return r.db.Delete(&entities.System{}, systemID).Error
}
