package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// EfficiencyMeasureRepository handles database operations for efficiency measures.
type EfficiencyMeasureRepository interface {
	CreateEfficiencyMeasure(measure *entities.EfficiencyMeasure) error
	ListEfficiencyMeasures() ([]entities.EfficiencyMeasure, error)
	FindEfficiencyMeasureByID(id uint) (*entities.EfficiencyMeasure, error)
	UpdateEfficiencyMeasureByID(id uint, updateData map[string]interface{}) error
}

type efficiencyMeasureRepository struct {
	db *gorm.DB
}

// NewEfficiencyMeasureRepository creates a new instance of EfficiencyMeasureRepository.
func NewEfficiencyMeasureRepository(db *gorm.DB) EfficiencyMeasureRepository {
	return &efficiencyMeasureRepository{db: db}
}

func (r *efficiencyMeasureRepository) CreateEfficiencyMeasure(measure *entities.EfficiencyMeasure) error {
	return r.db.Create(measure).Error
}

func (r *efficiencyMeasureRepository) ListEfficiencyMeasures() ([]entities.EfficiencyMeasure, error) {
	var measures []entities.EfficiencyMeasure
	err := r.db.Find(&measures).Error
	return measures, err
}

func (r *efficiencyMeasureRepository) FindEfficiencyMeasureByID(id uint) (*entities.EfficiencyMeasure, error) {
	var measure entities.EfficiencyMeasure
	err := r.db.First(&measure, id).Error
	if err != nil {
		return nil, err
	}
	return &measure, nil
}

func (r *efficiencyMeasureRepository) UpdateEfficiencyMeasureByID(id uint, updateData map[string]interface{}) error {
	return r.db.Model(&entities.EfficiencyMeasure{}).Where("id = ?", id).Updates(updateData).Error
}
