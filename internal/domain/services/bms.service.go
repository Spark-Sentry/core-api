package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// BmsService defines operations for Bms management.
type BmsService interface {
	CreateBms(req dto.CreateBmsRequest) (*entities.Bms, error)
	ListBms() ([]entities.Bms, error)
	GetBmsByID(id uint) (*entities.Bms, error)
	UpdateBmsByID(id uint, req dto.UpdateBmsRequest) error
	DeleteBmsByID(id uint) error
}

type bmsService struct {
	repo repository.BmsRepository
}

// NewBmsService creates a new instance of BmsService.
func NewBmsService(repo repository.BmsRepository) BmsService {
	return &bmsService{
		repo: repo,
	}
}

func (s *bmsService) CreateBms(req dto.CreateBmsRequest) (*entities.Bms, error) {
	bms := &entities.Bms{
		Type:      req.Type,
		IDPattern: req.IDPattern,
	}
	if err := s.repo.CreateBms(bms); err != nil {
		return nil, fmt.Errorf("failed to create Bms: %w", err)
	}
	return bms, nil
}

func (s *bmsService) ListBms() ([]entities.Bms, error) {
	return s.repo.ListBms()
}

func (s *bmsService) GetBmsByID(id uint) (*entities.Bms, error) {
	return s.repo.FindBmsByID(id)
}

func (s *bmsService) UpdateBmsByID(id uint, req dto.UpdateBmsRequest) error {
	updateData := map[string]interface{}{
		"type":       req.Type,
		"id_pattern": req.IDPattern,
	}
	if err := s.repo.UpdateBmsByID(id, updateData); err != nil {
		return fmt.Errorf("failed to update Bms: %w", err)
	}
	return nil
}

func (s *bmsService) DeleteBmsByID(id uint) error {
	return s.repo.DeleteBmsByID(id)
}
