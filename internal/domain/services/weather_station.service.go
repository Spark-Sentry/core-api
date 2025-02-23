package services

import (
	"core-api/internal/app/dto"
	"core-api/internal/domain/entities"
	"core-api/internal/infrastructure/repository"
	"fmt"
)

// WeatherStationService defines operations for weather station management.
type WeatherStationService interface {
	CreateWeatherStation(req dto.CreateWeatherStationRequest) (*entities.WeatherStation, error)
	ListWeatherStations() ([]entities.WeatherStation, error)
	GetWeatherStationByID(id uint) (*entities.WeatherStation, error)
	UpdateWeatherStationByID(id uint, req dto.UpdateWeatherStationRequest) error
	DeleteWeatherStationByID(id uint) error
}

type weatherStationService struct {
	repo repository.WeatherStationRepository
}

// NewWeatherStationService creates a new instance of WeatherStationService.
func NewWeatherStationService(repo repository.WeatherStationRepository) WeatherStationService {
	return &weatherStationService{
		repo: repo,
	}
}

func (s *weatherStationService) CreateWeatherStation(req dto.CreateWeatherStationRequest) (*entities.WeatherStation, error) {
	ws := &entities.WeatherStation{
		ClimateID: req.ClimateID,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
		Name:      req.Name,
	}
	if err := s.repo.CreateWeatherStation(ws); err != nil {
		return nil, fmt.Errorf("failed to create weather station: %w", err)
	}
	return ws, nil
}

func (s *weatherStationService) ListWeatherStations() ([]entities.WeatherStation, error) {
	return s.repo.ListWeatherStations()
}

func (s *weatherStationService) GetWeatherStationByID(id uint) (*entities.WeatherStation, error) {
	return s.repo.FindWeatherStationByID(id)
}

func (s *weatherStationService) UpdateWeatherStationByID(id uint, req dto.UpdateWeatherStationRequest) error {
	updateData := map[string]interface{}{
		"climate_id": req.ClimateID,
		"longitude":  req.Longitude,
		"latitude":   req.Latitude,
		"name":       req.Name,
	}
	if err := s.repo.UpdateWeatherStationByID(id, updateData); err != nil {
		return fmt.Errorf("failed to update weather station: %w", err)
	}
	return nil
}

func (s *weatherStationService) DeleteWeatherStationByID(id uint) error {
	return s.repo.DeleteWeatherStationByID(id)
}
