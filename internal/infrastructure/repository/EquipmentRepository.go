package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// EquipmentRepository handles database operations for equipments.
type EquipmentRepository interface {
	CreateEquipment(equip *entities.Equipment) error
	ListAllEquipments() ([]entities.Equipment, error)
	FindEquipmentByID(id uint) (*entities.Equipment, error)
	UpdateEquipmentByID(id uint, updateData map[string]interface{}) error
	DeleteEquipmentByID(id uint) error
}

type equipmentRepository struct {
	db *gorm.DB
}

// NewEquipmentRepository creates a new instance of EquipmentRepository.
func NewEquipmentRepository(db *gorm.DB) EquipmentRepository {
	return &equipmentRepository{db: db}
}

func (r *equipmentRepository) CreateEquipment(equip *entities.Equipment) error {
	return r.db.Create(equip).Error
}

func (r *equipmentRepository) ListAllEquipments() ([]entities.Equipment, error) {
	var equipments []entities.Equipment
	err := r.db.Preload("Parameters").Find(&equipments).Error
	return equipments, err
}

func (r *equipmentRepository) FindEquipmentByID(id uint) (*entities.Equipment, error) {
	var equip entities.Equipment
	err := r.db.Preload("Parameters").First(&equip, id).Error
	if err != nil {
		return nil, err
	}
	return &equip, nil
}

func (r *equipmentRepository) UpdateEquipmentByID(id uint, updateData map[string]interface{}) error {
	return r.db.Model(&entities.Equipment{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *equipmentRepository) DeleteEquipmentByID(id uint) error {
	return r.db.Delete(&entities.Equipment{}, id).Error
}
