package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// ContractorService defines operations related to contractors.
type ContractorService interface {
	CreateContractor(req dto.CreateContractorRequest) (*entities.Contractor, error)
	ListAllContractors() ([]entities.Contractor, error)
	GetContractorByID(id uint) (*entities.Contractor, error)
	UpdateContractor(id uint, req dto.UpdateContractorRequest) error
	DeleteContractor(id uint) error
}

type contractorService struct {
	contractorRepo repository.ContractorRepository
}

// NewContractorService creates a new instance of ContractorService.
func NewContractorService(contractorRepo repository.ContractorRepository) ContractorService {
	return &contractorService{
		contractorRepo: contractorRepo,
	}
}

func (s *contractorService) CreateContractor(req dto.CreateContractorRequest) (*entities.Contractor, error) {
	contractor := &entities.Contractor{
		Name:  req.Name,
		Phone: req.Phone,
		Email: req.Email,
	}
	if err := s.contractorRepo.CreateContractor(contractor); err != nil {
		return nil, fmt.Errorf("failed to create contractor: %w", err)
	}
	return contractor, nil
}

func (s *contractorService) ListAllContractors() ([]entities.Contractor, error) {
	return s.contractorRepo.ListAllContractors()
}

func (s *contractorService) GetContractorByID(id uint) (*entities.Contractor, error) {
	return s.contractorRepo.FindContractorByID(id)
}

func (s *contractorService) UpdateContractor(id uint, req dto.UpdateContractorRequest) error {
	updateData := map[string]interface{}{
		"name":  req.Name,
		"phone": req.Phone,
		"email": req.Email,
	}
	return s.contractorRepo.UpdateContractorByID(id, updateData)
}

func (s *contractorService) DeleteContractor(id uint) error {
	return s.contractorRepo.DeleteContractorByID(id)
}
