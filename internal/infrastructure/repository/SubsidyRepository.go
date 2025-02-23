package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// SubsidyRepository handles database operations for subsidies.
type SubsidyRepository interface {
	CreateSubsidy(subsidy *entities.Subsidy) error
	ListAllSubsidies() ([]entities.Subsidy, error)
	FindSubsidyByID(id uint) (*entities.Subsidy, error)
	UpdateSubsidyByID(id uint, updateData map[string]interface{}) error
	DeleteSubsidyByID(id uint) error
}

type subsidyRepository struct {
	db *gorm.DB
}

// NewSubsidyRepository creates a new instance of SubsidyRepository.
func NewSubsidyRepository(db *gorm.DB) SubsidyRepository {
	return &subsidyRepository{db: db}
}

func (r *subsidyRepository) CreateSubsidy(subsidy *entities.Subsidy) error {
	return r.db.Create(subsidy).Error
}

func (r *subsidyRepository) ListAllSubsidies() ([]entities.Subsidy, error) {
	var subsidies []entities.Subsidy
	err := r.db.Find(&subsidies).Error
	return subsidies, err
}

func (r *subsidyRepository) FindSubsidyByID(id uint) (*entities.Subsidy, error) {
	var subsidy entities.Subsidy
	err := r.db.First(&subsidy, id).Error
	if err != nil {
		return nil, err
	}
	return &subsidy, nil
}

func (r *subsidyRepository) UpdateSubsidyByID(id uint, updateData map[string]interface{}) error {
	return r.db.Model(&entities.Subsidy{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *subsidyRepository) DeleteSubsidyByID(id uint) error {
	return r.db.Delete(&entities.Subsidy{}, id).Error
}
