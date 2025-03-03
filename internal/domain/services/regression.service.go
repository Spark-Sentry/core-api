package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// RegressionService defines operations related to regressions.
type RegressionService interface {
	CreateRegression(req dto.CreateRegressionRequest) (*entities.Regression, error)
	ListRegressions() ([]entities.Regression, error)
	GetRegressionByID(id uint) (*entities.Regression, error)
	UpdateRegressionByID(id uint, req dto.UpdateRegressionRequest) error
	DeleteRegressionByID(id uint) error
}

type regressionService struct {
	repo      repository.RegressionRepository
	meterRepo repository.MeterRepository
}

// NewRegressionService creates a new instance of RegressionService.
func NewRegressionService(repo repository.RegressionRepository, meterRepo repository.MeterRepository) RegressionService {
	return &regressionService{
		repo:      repo,
		meterRepo: meterRepo,
	}
}

// CreateRegression creates a new regression along with its coefficients and meter associations.
func (s *regressionService) CreateRegression(req dto.CreateRegressionRequest) (*entities.Regression, error) {
	// Create regression entity
	regression := &entities.Regression{
		ProjectID: req.ProjectID,
		Name:      req.Name,
		Unit:      req.Unit,
	}
	// Create coefficients from the provided values.
	for _, value := range req.Coefficients {
		coef := entities.Coefficient{
			Value: value,
		}
		regression.Coefficients = append(regression.Coefficients, coef)
	}

	for _, meterID := range req.Meters {
		meter, err := s.meterRepo.FindMeterByID(meterID)
		if err != nil {
			return nil, fmt.Errorf("meter with id %s not found: %w", meterID, err)
		}
		regression.Meters = append(regression.Meters, *meter)
	}

	if err := s.repo.CreateRegression(regression); err != nil {
		return nil, fmt.Errorf("failed to create regression: %w", err)
	}
	return regression, nil
}

func (s *regressionService) ListRegressions() ([]entities.Regression, error) {
	return s.repo.ListAllRegressions()
}

func (s *regressionService) GetRegressionByID(id uint) (*entities.Regression, error) {
	return s.repo.FindRegressionByID(id)
}

func (s *regressionService) UpdateRegressionByID(id uint, req dto.UpdateRegressionRequest) error {
	updateData := map[string]interface{}{
		"name": req.Name,
		"unit": req.Unit,
		// Note: La mise à jour des coefficients et de la liste des meters nécessitera une gestion spécifique.
	}
	if err := s.repo.UpdateRegressionByID(id, updateData); err != nil {
		return fmt.Errorf("failed to update regression: %w", err)
	}
	// Pour simplifier, la mise à jour des coefficients et des meters n'est pas gérée ici.
	return nil
}

func (s *regressionService) DeleteRegressionByID(id uint) error {
	return s.repo.DeleteRegressionByID(id)
}
