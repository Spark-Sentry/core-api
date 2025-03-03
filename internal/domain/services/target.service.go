package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// TargetService defines operations related to targets.
type TargetService interface {
	CreateTarget(req dto.CreateTargetRequest) (*entities.Target, error)
	ListTargets() ([]entities.Target, error)
	GetTargetByID(id uint) (*entities.Target, error)
	UpdateTargetByID(id uint, req dto.UpdateTargetRequest) error
	DeleteTargetByID(id uint) error
}

type targetService struct {
	repo repository.TargetRepository
}

// NewTargetService creates a new instance of TargetService.
func NewTargetService(repo repository.TargetRepository) TargetService {
	return &targetService{
		repo: repo,
	}
}

func (s *targetService) CreateTarget(req dto.CreateTargetRequest) (*entities.Target, error) {
	target := &entities.Target{
		Name:                req.Name,
		Value:               req.Value,
		Type:                req.Type,
		ProjectID:           req.ProjectID,
		EfficiencyMeasureID: req.EfficiencyMeasureID,
	}
	if err := s.repo.CreateTarget(target); err != nil {
		return nil, fmt.Errorf("failed to create target: %w", err)
	}
	return target, nil
}

func (s *targetService) ListTargets() ([]entities.Target, error) {
	return s.repo.ListAllTargets()
}

func (s *targetService) GetTargetByID(id uint) (*entities.Target, error) {
	return s.repo.FindTargetByID(id)
}

func (s *targetService) UpdateTargetByID(id uint, req dto.UpdateTargetRequest) error {
	updateData := map[string]interface{}{
		"name":                  req.Name,
		"value":                 req.Value,
		"type":                  req.Type,
		"project_id":            req.ProjectID,
		"efficiency_measure_id": req.EfficiencyMeasureID,
	}
	if err := s.repo.UpdateTargetByID(id, updateData); err != nil {
		return fmt.Errorf("failed to update target: %w", err)
	}
	return nil
}

func (s *targetService) DeleteTargetByID(id uint) error {
	return s.repo.DeleteTargetByID(id)
}
