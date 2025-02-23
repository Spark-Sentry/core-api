package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// SubsidyService defines operations related to subsidies.
type SubsidyService interface {
	CreateSubsidy(req dto.CreateSubsidyRequest) (*entities.Subsidy, error)
	ListSubsidies() ([]entities.Subsidy, error)
	GetSubsidyByID(id uint) (*entities.Subsidy, error)
	UpdateSubsidyByID(id uint, req dto.UpdateSubsidyRequest) error
	DeleteSubsidyByID(id uint) error
}

type subsidyService struct {
	repo repository.SubsidyRepository
}

// NewSubsidyService creates a new instance of SubsidyService.
func NewSubsidyService(repo repository.SubsidyRepository) SubsidyService {
	return &subsidyService{
		repo: repo,
	}
}

func (s *subsidyService) CreateSubsidy(req dto.CreateSubsidyRequest) (*entities.Subsidy, error) {
	subsidy := &entities.Subsidy{
		Organisation:        req.Organisation,
		DateObtained:        req.DateObtained,
		Amount:              req.Amount,
		ProjectID:           req.ProjectID,
		EfficiencyMeasureID: req.EfficiencyMeasureID,
	}
	if err := s.repo.CreateSubsidy(subsidy); err != nil {
		return nil, fmt.Errorf("failed to create subsidy: %w", err)
	}
	return subsidy, nil
}

func (s *subsidyService) ListSubsidies() ([]entities.Subsidy, error) {
	return s.repo.ListAllSubsidies()
}

func (s *subsidyService) GetSubsidyByID(id uint) (*entities.Subsidy, error) {
	return s.repo.FindSubsidyByID(id)
}

func (s *subsidyService) UpdateSubsidyByID(id uint, req dto.UpdateSubsidyRequest) error {
	updateData := map[string]interface{}{
		"organisation":          req.Organisation,
		"date_obtained":         req.DateObtained,
		"amount":                req.Amount,
		"project_id":            req.ProjectID,
		"efficiency_measure_id": req.EfficiencyMeasureID,
	}
	if err := s.repo.UpdateSubsidyByID(id, updateData); err != nil {
		return fmt.Errorf("failed to update subsidy: %w", err)
	}
	return nil
}

func (s *subsidyService) DeleteSubsidyByID(id uint) error {
	return s.repo.DeleteSubsidyByID(id)
}
