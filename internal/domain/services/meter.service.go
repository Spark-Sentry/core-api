package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// MeterService defines operations related to meters.
type MeterService interface {
	CreateMeter(req dto.CreateMeterRequest) (*entities.Meter, error)
	ListAllMeters() ([]entities.Meter, error)
	GetMeterByID(id uint) (*entities.Meter, error)
	UpdateMeterByID(id uint, req dto.UpdateMeterRequest) error
	DeleteMeterByID(id uint) error
}

type meterService struct {
	repo repository.MeterRepository
}

// NewMeterService creates a new instance of MeterService.
func NewMeterService(repo repository.MeterRepository) MeterService {
	return &meterService{
		repo: repo,
	}
}

func (s *meterService) CreateMeter(req dto.CreateMeterRequest) (*entities.Meter, error) {
	meter := &entities.Meter{
		Supplier:   req.Supplier,
		BuildingID: req.BuildingID,
		Energy:     req.Energy,
		Name:       req.Name,
		MeterID:    req.MeterID,
	}
	if err := s.repo.CreateMeter(meter); err != nil {
		return nil, fmt.Errorf("failed to create meter: %w", err)
	}
	return meter, nil
}

func (s *meterService) ListAllMeters() ([]entities.Meter, error) {
	return s.repo.ListAllMeters()
}

func (s *meterService) GetMeterByID(id uint) (*entities.Meter, error) {
	return s.repo.FindMeterByID(id)
}

func (s *meterService) UpdateMeterByID(id uint, req dto.UpdateMeterRequest) error {
	updateData := map[string]interface{}{
		"supplier":    req.Supplier,
		"building_id": req.BuildingID,
		"energy":      req.Energy,
		"name":        req.Name,
		"meter_id":    req.MeterID,
	}
	if err := s.repo.UpdateMeterByID(id, updateData); err != nil {
		return fmt.Errorf("failed to update meter: %w", err)
	}
	return nil
}

func (s *meterService) DeleteMeterByID(id uint) error {
	return s.repo.DeleteMeterByID(id)
}
