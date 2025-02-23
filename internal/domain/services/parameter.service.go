package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// ParameterService defines operations related to parameters.
type ParameterService interface {
	CreateParameter(req dto.CreateParameterRequest) (*entities.Parameter, error)
	ListParameters() ([]entities.Parameter, error)
	GetParameterByID(id uint) (*entities.Parameter, error)
	UpdateParameterByID(id uint, req dto.UpdateParameterRequest) error
	DeleteParameterByID(id uint) error
}

type parameterService struct {
	repo repository.ParameterRepository
}

// NewParameterService creates a new instance of ParameterService.
func NewParameterService(repo repository.ParameterRepository) ParameterService {
	return &parameterService{
		repo: repo,
	}
}

func (s *parameterService) CreateParameter(req dto.CreateParameterRequest) (*entities.Parameter, error) {
	param := &entities.Parameter{
		Name:                req.Name,
		EquipmentID:         req.EquipmentID,
		BmsID:               req.BmsID,
		IDInBms:             req.IDInBms,
		PointType:           req.PointType,
		Unit:                req.Unit,
		EfficiencyMeasureID: req.EfficiencyMeasureID,
	}
	if err := s.repo.Create(param); err != nil {
		return nil, fmt.Errorf("failed to create parameter: %w", err)
	}
	return param, nil
}

func (s *parameterService) ListParameters() ([]entities.Parameter, error) {
	return s.repo.ListAll()
}

func (s *parameterService) GetParameterByID(id uint) (*entities.Parameter, error) {
	return s.repo.FindByID(id)
}

func (s *parameterService) UpdateParameterByID(id uint, req dto.UpdateParameterRequest) error {
	updateData := map[string]interface{}{
		"name":                  req.Name,
		"equipment_id":          req.EquipmentID,
		"bms_id":                req.BmsID,
		"id_in_bms":             req.IDInBms,
		"point_type":            req.PointType,
		"unit":                  req.Unit,
		"efficiency_measure_id": req.EfficiencyMeasureID,
	}
	if err := s.repo.Update(id, updateData); err != nil {
		return fmt.Errorf("failed to update parameter: %w", err)
	}
	return nil
}

func (s *parameterService) DeleteParameterByID(id uint) error {
	return s.repo.Delete(id)
}
