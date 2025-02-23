package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// TargetRepository handles database operations for targets.
type TargetRepository interface {
	CreateTarget(target *entities.Target) error
	ListAllTargets() ([]entities.Target, error)
	FindTargetByID(id uint) (*entities.Target, error)
	UpdateTargetByID(id uint, updateData map[string]interface{}) error
	DeleteTargetByID(id uint) error
}

type targetRepository struct {
	db *gorm.DB
}

// NewTargetRepository creates a new instance of TargetRepository.
func NewTargetRepository(db *gorm.DB) TargetRepository {
	return &targetRepository{db: db}
}

func (r *targetRepository) CreateTarget(target *entities.Target) error {
	return r.db.Create(target).Error
}

func (r *targetRepository) ListAllTargets() ([]entities.Target, error) {
	var targets []entities.Target
	err := r.db.Find(&targets).Error
	return targets, err
}

func (r *targetRepository) FindTargetByID(id uint) (*entities.Target, error) {
	var target entities.Target
	err := r.db.First(&target, id).Error
	if err != nil {
		return nil, err
	}
	return &target, nil
}

func (r *targetRepository) UpdateTargetByID(id uint, updateData map[string]interface{}) error {
	return r.db.Model(&entities.Target{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *targetRepository) DeleteTargetByID(id uint) error {
	return r.db.Delete(&entities.Target{}, id).Error
}
