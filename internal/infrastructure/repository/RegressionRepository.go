package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// RegressionRepository handles database operations for regressions.
type RegressionRepository interface {
	CreateRegression(regression *entities.Regression) error
	ListAllRegressions() ([]entities.Regression, error)
	FindRegressionByID(id uint) (*entities.Regression, error)
	UpdateRegressionByID(id uint, updateData map[string]interface{}) error
	DeleteRegressionByID(id uint) error
}

type regressionRepository struct {
	db *gorm.DB
}

// NewRegressionRepository creates a new instance of RegressionRepository.
func NewRegressionRepository(db *gorm.DB) RegressionRepository {
	return &regressionRepository{db: db}
}

func (r *regressionRepository) CreateRegression(regression *entities.Regression) error {
	return r.db.Create(regression).Error
}

func (r *regressionRepository) ListAllRegressions() ([]entities.Regression, error) {
	var regressions []entities.Regression
	err := r.db.Preload("Coefficients").Preload("Meters").Find(&regressions).Error
	return regressions, err
}

func (r *regressionRepository) FindRegressionByID(id uint) (*entities.Regression, error) {
	var regression entities.Regression
	err := r.db.Preload("Coefficients").Preload("Meters").First(&regression, id).Error
	if err != nil {
		return nil, err
	}
	return &regression, nil
}

func (r *regressionRepository) UpdateRegressionByID(id uint, updateData map[string]interface{}) error {
	return r.db.Model(&entities.Regression{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *regressionRepository) DeleteRegressionByID(id uint) error {
	return r.db.Delete(&entities.Regression{}, id).Error
}
