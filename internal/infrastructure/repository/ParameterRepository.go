package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

type ParameterRepository interface {
	Create(param *entities.Parameter) error
	ListAll() ([]entities.Parameter, error)
	FindByID(id uint) (*entities.Parameter, error)
	Update(id uint, updateData map[string]interface{}) error
	Delete(id uint) error
	DB() *gorm.DB
}

type parameterRepository struct {
	db *gorm.DB
}

func NewParameterRepository(db *gorm.DB) ParameterRepository {
	return &parameterRepository{db: db}
}

func (r *parameterRepository) Create(param *entities.Parameter) error {
	return r.db.Create(param).Error
}

func (r *parameterRepository) ListAll() ([]entities.Parameter, error) {
	var params []entities.Parameter
	err := r.db.Preload("EfficiencyMeasures").Find(&params).Error
	return params, err
}

func (r *parameterRepository) FindByID(id uint) (*entities.Parameter, error) {
	var param entities.Parameter
	err := r.db.Preload("EfficiencyMeasures").First(&param, id).Error
	if err != nil {
		return nil, err
	}
	return &param, nil
}

func (r *parameterRepository) Update(id uint, updateData map[string]interface{}) error {
	return r.db.Model(&entities.Parameter{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *parameterRepository) Delete(id uint) error {
	return r.db.Delete(&entities.Parameter{}, id).Error
}

func (r *parameterRepository) DB() *gorm.DB {
	return r.db
}
