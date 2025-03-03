package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

type ParameterService interface {
	CreateParameter(req dto.CreateParameterRequest) (*entities.Parameter, error)
	ListParameters() ([]entities.Parameter, error)
	GetParameterByID(id uint) (*entities.Parameter, error)
	UpdateParameterByID(id uint, req dto.UpdateParameterRequest) error
	DeleteParameterByID(id uint) error
}

type parameterService struct {
	repo           repository.ParameterRepository
	efficiencyRepo repository.EfficiencyMeasureRepository
}

func NewParameterService(repo repository.ParameterRepository, efficiencyRepo repository.EfficiencyMeasureRepository) ParameterService {
	return &parameterService{
		repo:           repo,
		efficiencyRepo: efficiencyRepo,
	}
}

func (s *parameterService) CreateParameter(req dto.CreateParameterRequest) (*entities.Parameter, error) {
	// Create the base Parameter record
	param := &entities.Parameter{
		Name:        req.Name,
		EquipmentID: req.EquipmentID,
		BmsID:       req.BmsID,
		IDInBms:     req.IDInBms,
		PointType:   req.PointType,
		Unit:        req.Unit,
	}
	if err := s.repo.Create(param); err != nil {
		return nil, fmt.Errorf("failed to create parameter: %w", err)
	}

	// Retrieve EfficiencyMeasure entities for the provided IDs.
	var effMeasures []entities.EfficiencyMeasure
	for _, id := range req.EfficiencyMeasureIDs {
		em, err := s.efficiencyRepo.FindEfficiencyMeasureByID(id)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve efficiency measure with id %d: %w", id, err)
		}
		if em == nil {
			return nil, fmt.Errorf("efficiency measure with id %d not found", id)
		}
		effMeasures = append(effMeasures, *em)
	}

	// Associate the EfficiencyMeasures using GORM's association mode.
	if err := s.repo.DB().Model(param).Association("EfficiencyMeasures").Replace(effMeasures); err != nil {
		return nil, fmt.Errorf("failed to associate efficiency measures: %w", err)
	}

	// Optionally, reload the parameter with its associations.
	updatedParam, err := s.repo.FindByID(param.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to reload parameter: %w", err)
	}
	return updatedParam, nil
}

func (s *parameterService) ListParameters() ([]entities.Parameter, error) {
	return s.repo.ListAll()
}

func (s *parameterService) GetParameterByID(id uint) (*entities.Parameter, error) {
	return s.repo.FindByID(id)
}

func (s *parameterService) UpdateParameterByID(id uint, req dto.UpdateParameterRequest) error {
	// Update base fields.
	updateData := map[string]interface{}{
		"name":         req.Name,
		"equipment_id": req.EquipmentID,
		"bms_id":       req.BmsID,
		"id_in_bms":    req.IDInBms,
		"point_type":   req.PointType,
		"unit":         req.Unit,
	}
	if err := s.repo.Update(id, updateData); err != nil {
		return fmt.Errorf("failed to update parameter: %w", err)
	}

	// Retrieve the updated parameter.
	param, err := s.repo.FindByID(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve updated parameter: %w", err)
	}

	// Retrieve EfficiencyMeasure entities for the provided IDs.
	var effMeasures []entities.EfficiencyMeasure
	for _, id := range req.EfficiencyMeasureIDs {
		em, err := s.efficiencyRepo.FindEfficiencyMeasureByID(id)
		if err != nil {
			return fmt.Errorf("failed to retrieve efficiency measure with id %d: %w", id, err)
		}
		if em == nil {
			return fmt.Errorf("efficiency measure with id %d not found", id)
		}
		effMeasures = append(effMeasures, *em)
	}

	// Update the associations.
	if err := s.repo.DB().Model(param).Association("EfficiencyMeasures").Replace(effMeasures); err != nil {
		return fmt.Errorf("failed to update efficiency measure associations: %w", err)
	}

	return nil
}

func (s *parameterService) DeleteParameterByID(id uint) error {
	return s.repo.Delete(id)
}
