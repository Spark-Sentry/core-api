package repository

import (
	"core-api/internal/domain/entities"
	"gorm.io/gorm"
)

// WeatherStationRepository handles database operations for weather stations.
type WeatherStationRepository interface {
	CreateWeatherStation(ws *entities.WeatherStation) error
	ListWeatherStations() ([]entities.WeatherStation, error)
	FindWeatherStationByID(id uint) (*entities.WeatherStation, error)
	UpdateWeatherStationByID(id uint, updateData map[string]interface{}) error
	DeleteWeatherStationByID(id uint) error
}

type weatherStationRepository struct {
	db *gorm.DB
}

// NewWeatherStationRepository creates a new instance of WeatherStationRepository.
func NewWeatherStationRepository(db *gorm.DB) WeatherStationRepository {
	return &weatherStationRepository{db: db}
}

func (r *weatherStationRepository) CreateWeatherStation(ws *entities.WeatherStation) error {
	return r.db.Create(ws).Error
}

func (r *weatherStationRepository) ListWeatherStations() ([]entities.WeatherStation, error) {
	var wsList []entities.WeatherStation
	err := r.db.Find(&wsList).Error
	return wsList, err
}

func (r *weatherStationRepository) FindWeatherStationByID(id uint) (*entities.WeatherStation, error) {
	var ws entities.WeatherStation
	err := r.db.First(&ws, id).Error
	if err != nil {
		return nil, err
	}
	return &ws, nil
}

func (r *weatherStationRepository) UpdateWeatherStationByID(id uint, updateData map[string]interface{}) error {
	return r.db.Model(&entities.WeatherStation{}).Where("id = ?", id).Updates(updateData).Error
}

func (r *weatherStationRepository) DeleteWeatherStationByID(id uint) error {
	return r.db.Delete(&entities.WeatherStation{}, id).Error
}
