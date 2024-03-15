package repository

import (
	"core-api/internal/app/dto"
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

// FindBySystemID retrieves all equipments associated with a specific system.
func (r *EquipmentRepository) FindBySystemID(systemID uint) ([]entities.Equipment, error) {
	var equipments []entities.Equipment
	err := r.db.Preload("System").Where("system_id = ?", systemID).Find(&equipments).Error
	return equipments, err
}

// UpdateEquipment updates an existing piece of equipment with new details.
func (r *EquipmentRepository) UpdateEquipment(equipmentID uint, updateDTO dto.EquipmentUpdateDTO) error {
	return r.db.Model(&entities.Equipment{}).Where("id = ?", equipmentID).Updates(entities.Equipment{Tag: updateDTO.Tag, Description: updateDTO.Description}).Error
}

// DeleteEquipment deletes an existing piece of equipment.
func (r *EquipmentRepository) DeleteEquipment(equipmentID uint) error {
	return r.db.Delete(&entities.Equipment{}, equipmentID).Error
}
