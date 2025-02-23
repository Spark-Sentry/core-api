package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// IndependantVariableRepository handles database operations for independant variables.
type IndependantVariableRepository interface {
	Create(iv *entities.IndependantVariable) error
	ListAll() ([]entities.IndependantVariable, error)
	FindByID(id uint) (*entities.IndependantVariable, error)
	Update(id uint, updateData map[string]interface{}) error
	Delete(id uint) error
}

type independantVariableRepository struct {
	db *gorm.DB
}

// NewIndependantVariableRepository creates a new instance of IndependantVariableRepository.
func NewIndependantVariableRepository(db *gorm.DB) IndependantVariableRepository {
	return &independantVariableRepository{db: db}
}

func (r *independantVariableRepository) Create(iv *entities.IndependantVariable) error {
	return r.db.Create(iv).Error
}

func (r *independantVariableRepository) ListAll() ([]entities.IndependantVariable, error) {
	var ivs []entities.IndependantVariable
	err := r.db.Find(&ivs).Error
	return ivs, err
}

func (r *independantVariableRepository) FindByID(id uint) (*entities.IndependantVariable, error) {
	var iv entities.IndependantVariable
	err := r.db.First(&iv, id).Error
	if err != nil {
		return nil, err
	}
	return &iv, nil
}

func (r *independantVariableRepository) Update(id uint, updateData map[string]interface{}) error {
	return r.db.Model(&entities.IndependantVariable{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *independantVariableRepository) Delete(id uint) error {
	return r.db.Delete(&entities.IndependantVariable{}, id).Error
}
