package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// ContractorRepository handles database operations for contractors.
type ContractorRepository interface {
	CreateContractor(contractor *entities.Contractor) error
	ListAllContractors() ([]entities.Contractor, error)
	FindContractorByID(id uint) (*entities.Contractor, error)
	UpdateContractorByID(id uint, updateData map[string]interface{}) error
	DeleteContractorByID(id uint) error
}

type contractorRepository struct {
	db *gorm.DB
}

// NewContractorRepository creates a new instance of ContractorRepository.
func NewContractorRepository(db *gorm.DB) ContractorRepository {
	return &contractorRepository{db: db}
}

func (r *contractorRepository) CreateContractor(contractor *entities.Contractor) error {
	return r.db.Create(contractor).Error
}

func (r *contractorRepository) ListAllContractors() ([]entities.Contractor, error) {
	var contractors []entities.Contractor
	err := r.db.Find(&contractors).Error
	return contractors, err
}

func (r *contractorRepository) FindContractorByID(id uint) (*entities.Contractor, error) {
	var contractor entities.Contractor
	err := r.db.First(&contractor, id).Error
	if err != nil {
		return nil, err
	}
	return &contractor, nil
}

func (r *contractorRepository) UpdateContractorByID(id uint, updateData map[string]interface{}) error {
	return r.db.Model(&entities.Contractor{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *contractorRepository) DeleteContractorByID(id uint) error {
	return r.db.Delete(&entities.Contractor{}, id).Error
}
