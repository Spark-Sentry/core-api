package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// AreaRepository interface for CRUD operations on Areas.
type AreaRepository interface {
	Save(area *entities.Area) (*entities.Area, error)
	FindByBuildingID(buildingID uint) ([]entities.Area, error)
	UpdateArea(area entities.Area) error
	DeleteArea(areaID uint) error
}

type areaRepository struct {
	db *gorm.DB
}

// NewAreaRepository creates a new instance of AreaRepository.
func NewAreaRepository(db *gorm.DB) AreaRepository {
	return &areaRepository{db: db}
}

// Save inserts a new Area into the database.
func (r *areaRepository) Save(area *entities.Area) (*entities.Area, error) {
	if err := r.db.Create(area).Error; err != nil {
		return nil, err
	}
	return area, nil
}

// FindByBuildingID retrieves areas by their associated building ID
func (r *areaRepository) FindByBuildingID(buildingID uint) ([]entities.Area, error) {
	var areas []entities.Area
	err := r.db.Where("building_id = ?", buildingID).Find(&areas).Error
	return areas, err
}

// UpdateArea updates an existing area with new details.
func (r *areaRepository) UpdateArea(area entities.Area) error {
	return r.db.Model(&entities.Area{}).Where("id = ?", area.ID).Updates(area).Error
}

// DeleteArea deletes an existing area.
func (r *areaRepository) DeleteArea(areaID uint) error {
	return r.db.Delete(&entities.Area{}, areaID).Error
}
