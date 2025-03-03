package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// EfficiencyMeasureService defines operations related to efficiency measures.
type EfficiencyMeasureService interface {
	CreateEfficiencyMeasure(measure *entities.EfficiencyMeasure) error
	ListEfficiencyMeasures() ([]entities.EfficiencyMeasure, error)
	GetEfficiencyMeasureByID(id uint) (*entities.EfficiencyMeasure, error)
	UpdateEfficiencyMeasureByID(id uint, req dto.UpdateEfficiencyMeasureRequest) error
}

type efficiencyMeasureService struct {
	repo repository.EfficiencyMeasureRepository
}

// NewEfficiencyMeasureService creates a new instance of EfficiencyMeasureService.
func NewEfficiencyMeasureService(repo repository.EfficiencyMeasureRepository) EfficiencyMeasureService {
	return &efficiencyMeasureService{
		repo: repo,
	}
}

func (s *efficiencyMeasureService) CreateEfficiencyMeasure(measure *entities.EfficiencyMeasure) error {
	if measure.Name == "" || measure.ProjectID == 0 {
		return fmt.Errorf("invalid efficiency measure data")
	}
	return s.repo.CreateEfficiencyMeasure(measure)
}

func (s *efficiencyMeasureService) ListEfficiencyMeasures() ([]entities.EfficiencyMeasure, error) {
	return s.repo.ListEfficiencyMeasures()
}

func (s *efficiencyMeasureService) GetEfficiencyMeasureByID(id uint) (*entities.EfficiencyMeasure, error) {
	return s.repo.FindEfficiencyMeasureByID(id)
}

// UpdateEfficiencyMeasureByID updates an efficiency measure by its ID using a map.
func (s *efficiencyMeasureService) UpdateEfficiencyMeasureByID(id uint, req dto.UpdateEfficiencyMeasureRequest) error {
	updateData := map[string]interface{}{
		"name":             req.Name,
		"efficiency_cost":  req.EfficiencyCost,
		"maintenance_cost": req.MaintenanceCost,
	}
	if err := s.repo.UpdateEfficiencyMeasureByID(id, updateData); err != nil {
		return fmt.Errorf("failed to update efficiency measure: %w", err)
	}
	return nil
}
