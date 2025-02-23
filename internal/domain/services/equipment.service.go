package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// EquipmentService defines operations related to equipments.
type EquipmentService interface {
	CreateEquipment(req dto.CreateEquipmentRequest) (*entities.Equipment, error)
	ListAllEquipments() ([]entities.Equipment, error)
	GetEquipmentByID(id uint) (*entities.Equipment, error)
	UpdateEquipmentByID(id uint, req dto.UpdateEquipmentRequest) error
	DeleteEquipmentByID(id uint) error
}

type equipmentService struct {
	repo repository.EquipmentRepository
}

// NewEquipmentService creates a new instance of EquipmentService.
func NewEquipmentService(repo repository.EquipmentRepository) EquipmentService {
	return &equipmentService{
		repo: repo,
	}
}

func (s *equipmentService) CreateEquipment(req dto.CreateEquipmentRequest) (*entities.Equipment, error) {
	equip := &entities.Equipment{
		Name: req.Name,
		Tag:  req.Tag,
	}
	if err := s.repo.CreateEquipment(equip); err != nil {
		return nil, fmt.Errorf("failed to create equipment: %w", err)
	}
	return equip, nil
}

func (s *equipmentService) ListAllEquipments() ([]entities.Equipment, error) {
	return s.repo.ListAllEquipments()
}

func (s *equipmentService) GetEquipmentByID(id uint) (*entities.Equipment, error) {
	return s.repo.FindEquipmentByID(id)
}

func (s *equipmentService) UpdateEquipmentByID(id uint, req dto.UpdateEquipmentRequest) error {
	updateData := map[string]interface{}{
		"name": req.Name,
		"tag":  req.Tag,
	}
	if err := s.repo.UpdateEquipmentByID(id, updateData); err != nil {
		return fmt.Errorf("failed to update equipment: %w", err)
	}
	return nil
}

func (s *equipmentService) DeleteEquipmentByID(id uint) error {
	return s.repo.DeleteEquipmentByID(id)
}
