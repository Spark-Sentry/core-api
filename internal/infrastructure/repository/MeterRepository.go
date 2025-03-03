package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// MeterRepository handles database operations for meters.
type MeterRepository interface {
	CreateMeter(meter *entities.Meter) error
	ListAllMeters() ([]entities.Meter, error)
	FindMeterByID(id uint) (*entities.Meter, error)
	UpdateMeterByID(id uint, updateData map[string]interface{}) error
	DeleteMeterByID(id uint) error
}

type meterRepository struct {
	db *gorm.DB
}

// NewMeterRepository creates a new instance of MeterRepository.
func NewMeterRepository(db *gorm.DB) MeterRepository {
	return &meterRepository{db: db}
}

func (r *meterRepository) CreateMeter(meter *entities.Meter) error {
	return r.db.Create(meter).Error
}

func (r *meterRepository) ListAllMeters() ([]entities.Meter, error) {
	var meters []entities.Meter
	err := r.db.Find(&meters).Error
	return meters, err
}

func (r *meterRepository) FindMeterByID(id uint) (*entities.Meter, error) {
	var meter entities.Meter
	err := r.db.First(&meter, id).Error
	if err != nil {
		return nil, err
	}
	return &meter, nil
}

func (r *meterRepository) UpdateMeterByID(id uint, updateData map[string]interface{}) error {
	return r.db.Model(&entities.Meter{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *meterRepository) DeleteMeterByID(id uint) error {
	return r.db.Delete(&entities.Meter{}, id).Error
}
