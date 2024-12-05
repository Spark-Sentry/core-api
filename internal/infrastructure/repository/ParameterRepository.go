package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// ParameterRepository defines the methods to interact with the Parameter table.
type ParameterRepository struct {
	db *gorm.DB
}

// NewParameterRepository creates a new instance of ParameterRepository.
func NewParameterRepository(db *gorm.DB) *ParameterRepository {
	return &ParameterRepository{db: db}
}

// Create creates a new Parameter in the database.
func (r *ParameterRepository) Create(parameter *entities.Parameter) error {
	return r.db.Create(parameter).Error
}

// FindByID retrieves a Parameter by its ID.
func (r *ParameterRepository) FindByID(id uint) (*entities.Parameter, error) {
	var parameter entities.Parameter
	if err := r.db.Preload("Equipment").First(&parameter, id).Error; err != nil {
		return nil, err
	}
	return &parameter, nil
}

// FindAll retrieves all Parameters from the database.
func (r *ParameterRepository) FindAll() ([]entities.Parameter, error) {
	var parameters []entities.Parameter
	if err := r.db.Preload("Equipment").Find(&parameters).Error; err != nil {
		return nil, err
	}
	return parameters, nil
}

// Update updates an existing Parameter in the database.
func (r *ParameterRepository) Update(parameter *entities.Parameter) error {
	return r.db.Save(parameter).Error
}

// Delete deletes a Parameter from the database by ID.
func (r *ParameterRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Parameter{}, id).Error
}

// FindByEquipmentID retrieves all Parameters associated with a specific Equipment ID.
func (r *ParameterRepository) FindByEquipmentID(equipmentID uint) ([]entities.Parameter, error) {
	var parameters []entities.Parameter
	if err := r.db.Where("equipment_id = ?", equipmentID).Find(&parameters).Error; err != nil {
		return nil, err
	}
	return parameters, nil
}
