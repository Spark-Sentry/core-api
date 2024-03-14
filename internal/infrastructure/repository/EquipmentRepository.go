package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// EquipmentRepository handles database operations for Equipment.
type EquipmentRepository struct {
	db *gorm.DB
}

func NewEquipmentRepository(db *gorm.DB) *EquipmentRepository {
	return &EquipmentRepository{db: db}
}

// AddEquipment adds new Equipment to the database.
func (r *EquipmentRepository) AddEquipment(equipment *entities.Equipment) error {
	return r.db.Create(equipment).Error
}
