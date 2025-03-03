package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// IndependantVariableService defines operations for independant variables.
type IndependantVariableService interface {
	CreateIndependantVariable(req dto.CreateIndependantVariableRequest) (*entities.IndependantVariable, error)
	ListIndependantVariables() ([]entities.IndependantVariable, error)
	GetIndependantVariableByID(id uint) (*entities.IndependantVariable, error)
	UpdateIndependantVariableByID(id uint, req dto.UpdateIndependantVariableRequest) error
	DeleteIndependantVariableByID(id uint) error
}

type independantVariableService struct {
	repo repository.IndependantVariableRepository
}

// NewIndependantVariableService creates a new instance of IndependantVariableService.
func NewIndependantVariableService(repo repository.IndependantVariableRepository) IndependantVariableService {
	return &independantVariableService{
		repo: repo,
	}
}

func (s *independantVariableService) CreateIndependantVariable(req dto.CreateIndependantVariableRequest) (*entities.IndependantVariable, error) {
	iv := &entities.IndependantVariable{
		Name:      req.Name,
		DateStart: req.DateStart,
		DateStop:  req.DateStop,
		Value:     req.Value,
		Source:    req.Source,
	}
	if err := s.repo.Create(iv); err != nil {
		return nil, fmt.Errorf("failed to create independant variable: %w", err)
	}
	return iv, nil
}

func (s *independantVariableService) ListIndependantVariables() ([]entities.IndependantVariable, error) {
	return s.repo.ListAll()
}

func (s *independantVariableService) GetIndependantVariableByID(id uint) (*entities.IndependantVariable, error) {
	return s.repo.FindByID(id)
}

func (s *independantVariableService) UpdateIndependantVariableByID(id uint, req dto.UpdateIndependantVariableRequest) error {
	updateData := map[string]interface{}{
		"name":       req.Name,
		"date_start": req.DateStart,
		"date_stop":  req.DateStop,
		"value":      req.Value,
		"source":     req.Source,
	}
	if err := s.repo.Update(id, updateData); err != nil {
		return fmt.Errorf("failed to update independant variable: %w", err)
	}
	return nil
}

func (s *independantVariableService) DeleteIndependantVariableByID(id uint) error {
	return s.repo.Delete(id)
}
